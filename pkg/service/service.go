package service

import (
	"context"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"strings"

	"github.com/demophoon/shrls/pkg/config"
	"github.com/demophoon/shrls/server"
	shrls "github.com/demophoon/shrls/server"
	pb "github.com/demophoon/shrls/server/gen"
	gw "github.com/demophoon/shrls/server/gen/gateway"
	"github.com/demophoon/shrls/service"
	boltstate "github.com/demophoon/shrls/state/boltdb"
	mongostate "github.com/demophoon/shrls/state/mongo"

	directorystate "github.com/demophoon/shrls/storage/directory"
	"github.com/demophoon/shrls/ui"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/selector"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type ShrlsService struct {
	server  shrls.ServerHandler
	State   shrls.ServerState
	storage shrls.ServerStorage
	config  *config.Config
}

func (s *ShrlsService) checkBasicAuthOk(username string, password string) bool {
	usernameHash := sha256.Sum256([]byte(username))
	passwordHash := sha256.Sum256([]byte(password))
	configUsernameHash := sha256.Sum256([]byte(s.config.AuthBackend.Basic.Username))
	configPasswordHash := sha256.Sum256([]byte(s.config.AuthBackend.Basic.Password))

	usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], configUsernameHash[:]) == 1)
	passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], configPasswordHash[:]) == 1)

	if usernameMatch && passwordMatch {
		return true
	}
	return false
}

func (s *ShrlsService) BasicAuthHandler(next http.HandlerFunc) http.HandlerFunc {
	// Backend auth not configured
	if s.config.AuthBackend == nil {
		return next
	}

	// Basic auth configured
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if ok {
			if s.checkBasicAuthOk(username, password) {
				next.ServeHTTP(w, r)
				return
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="shrls admin", charset="UTF-8"`)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}

func (s *ShrlsService) GRPCAuth(ctx context.Context) (context.Context, error) {
	unauthenticated_error := status.Error(codes.Unauthenticated, "Unauthenticated")

	// Backend auth not configured
	if s.config.AuthBackend == nil {
		return ctx, nil
	}

	token, err := auth.AuthFromMD(ctx, "basic")
	if err != nil {
		log.Errorf("Unautnenticated: %s", err)
		return nil, unauthenticated_error
	}

	c, err := base64.StdEncoding.DecodeString(token[:])
	if err != nil {
		log.Errorf("Error decoding token: %s", err)
		return nil, unauthenticated_error
	}
	cs := string(c)
	username, password, ok := strings.Cut(cs, ":")
	if !ok {
		log.Errorf("Invalid basic auth string: %s", err)
		return nil, unauthenticated_error
	}
	if s.checkBasicAuthOk(username, password) {
		return ctx, nil
	}

	return nil, unauthenticated_error
}

func (s *ShrlsService) Run() error {
	log.Debug("Setting up gRPC server")
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			selector.UnaryServerInterceptor(
				auth.UnaryServerInterceptor(s.GRPCAuth),
				selector.MatchFunc(func(ctx context.Context, meta interceptors.CallMeta) bool {
					return true
				}),
			),
		),
	)
	reflection.Register(grpcServer)

	log.Debug("Registering SHRLS")
	pb.RegisterShrlsServer(grpcServer, s.server)
	log.Debug("Registering Upload Service")
	pb.RegisterFileUploadServer(grpcServer, s.server)

	ctx := context.Background()
	grpcAddr := fmt.Sprintf("localhost:%d", s.config.Port)
	grpcOpts := []grpc.DialOption{grpc.WithInsecure()}

	log.Debug("Registering SHRLS HTTP endpoints")
	mux := runtime.NewServeMux()
	err := gw.RegisterShrlsHandlerFromEndpoint(ctx, mux, grpcAddr, grpcOpts)
	if err != nil {
		return err
	}
	err = gw.RegisterFileUploadHandlerFromEndpoint(ctx, mux, grpcAddr, grpcOpts)

	log.Debug("Registering UI")
	sub, err := fs.Sub(ui.Content, "dist")
	if err != nil {
		panic(err)
	}
	fs := http.FileServer(http.FS(sub))

	log.Debug("Adding Handlers")
	main := http.NewServeMux()
	main.Handle("/admin/", s.BasicAuthHandler(http.StripPrefix("/admin/", fs).ServeHTTP))
	main.Handle("/v1/", s.BasicAuthHandler(mux.ServeHTTP))
	main.HandleFunc("/", s.Redirect)

	server := http.Server{
		Handler: main,
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.Port))
	if err != nil {
		return err
	}

	m := cmux.New(l)
	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.Match(cmux.HTTP2())

	log.Debug("Starting servers")
	go server.Serve(httpL)
	go grpcServer.Serve(grpcL)

	log.Infof("Listening on :%d", s.config.Port)
	return m.Serve()
}

func (s *ShrlsService) SetState(state shrls.ServerState) {
	s.State = state
}

func (s *ShrlsService) SetServer(server shrls.ServerHandler) {
	s.server = server
}

func (s *ShrlsService) SetStorage(storage shrls.ServerStorage) {
	s.storage = storage
}

func (s *ShrlsService) SetConfig(config *config.Config) {
	s.config = config
}

// This function is responsible for returning a concrete implementation of the
// Shrls service with a mongodb state backend. Other backends can be setup by
// manually configuring the ShrlsService{} type itself.
func New(config *config.Config) ShrlsService {
	s := ShrlsService{}

	log.Debug("Initalizing config")
	s.SetConfig(config)

	// Set ServerStorage
	log.Debug("Configuring storage")
	storage := directorystate.New(config)
	s.SetStorage(storage)

	// Set ServerState
	log.Debug("Configuring backend state")
	if config.StateBackend == nil {
		log.Fatal("State backend is undefined")
	}
	var state server.ServerState
	if config.StateBackend.Bolt != nil {
		state = boltstate.New(*config)
		s.SetState(state)
	}
	if config.StateBackend.Mongo != nil {
		state = mongostate.New(*config)
		s.SetState(state)
	}

	// Set Server Implementation
	log.Debug("Adding server implementation")
	impl := service.New(config)
	impl.SetState(state)
	impl.SetStorage(storage)
	s.SetServer(impl)

	log.Debug("Server initialized")

	return s
}
