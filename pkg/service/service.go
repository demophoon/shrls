package service

import (
	"context"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"strings"
	"sync"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
	shrls "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
	gw "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen/gateway"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/service"
	mongostate "gitlab.cascadia.demophoon.com/demophoon/go-shrls/state/mongo"
	directorystate "gitlab.cascadia.demophoon.com/demophoon/go-shrls/storage/directory"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/ui"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type ShrlsService struct {
	server  *shrls.ServerHandler
	State   *shrls.ServerState
	storage *shrls.ServerStorage
	config  *config.Config
}

func (s *ShrlsService) NewClient() shrls.ServerState {
	return *s.State
}

func (s *ShrlsService) Run() error {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.startGRpcServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.startHttpServer()
	}()

	wg.Wait()
	return nil
}

func (s *ShrlsService) Redirect(w http.ResponseWriter, r *http.Request) {
	shrl := r.URL.Path
	shrl = strings.TrimPrefix(shrl, "/")

	parts := strings.Split(shrl, ".")
	ref := &pb.Ref_ShortURL{
		Ref: &pb.Ref_ShortURL_Alias{
			Alias: parts[0],
		},
	}

	ctx := context.Background()
	redirect, err := s.NewClient().GetShrl(ctx, ref)
	if err != nil {
		if s.config.DefaultRedirect != "" {
			http.Redirect(w, r, s.config.DefaultRedirect, http.StatusTemporaryRedirect)
			log.Error("Unable to resolve Shrl. ", ref)
		} else {
			http.Error(w, fmt.Sprintf("Unable to fetch Shrl: %q. %s", ref, err), http.StatusNotFound)
		}
		return
	}

	switch redirect.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		http.Redirect(w, r, redirect.Content.GetUrl().Url, http.StatusTemporaryRedirect)

	case *pb.ExpandedURL_File:
		file := redirect.Content.GetFile()
		w.Write(file)

	case *pb.ExpandedURL_Snippet:
		w.Write([]byte(redirect.Content.GetSnippet().Body))
	}

	log.Debug(redirect)
}

func (s ShrlsService) startHttpServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	grpcPort := viper.GetInt("grpc_port")
	err := gw.RegisterShrlsHandlerFromEndpoint(ctx, mux, fmt.Sprintf("localhost:%d", grpcPort), opts)
	if err != nil {
		return err
	}

	// Add static file assets
	sub, err := fs.Sub(ui.Content, "dist")
	if err != nil {
		panic(err)
	}
	fs := http.FileServer(http.FS(sub))

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Info(fmt.Sprintf("Starting HTTP Server: %d", viper.GetInt("port")))

	main := http.NewServeMux()
	main.Handle("/admin/", http.StripPrefix("/admin/", fs))
	main.Handle("/v1/", mux)
	main.HandleFunc("/", s.Redirect)

	return http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), main)
}

func (s ShrlsService) startGRpcServer() error {
	port := viper.GetInt("grpc_port")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterShrlsServer(grpcServer, *s.server)

	reflection.Register(grpcServer)

	log.Info(fmt.Sprintf("Starting gRPC Server: %d", port))
	grpcServer.Serve(lis)

	return nil
}

func (s *ShrlsService) SetState(state shrls.ServerState) {
	s.State = &state
}

func (s *ShrlsService) SetServer(server shrls.ServerHandler) {
	s.server = &server
}

func (s *ShrlsService) SetStorage(storage shrls.ServerStorage) {
	s.storage = &storage
}

func (s *ShrlsService) SetConfig(config *config.Config) {
	s.config = config
}

// This function is responsible for returning a concrete implementation of the
// Shrls service with a mongodb state backend. Other backends can be setup by
// manually configuring the ShrlsService{} type itself.
func New(config *config.Config) ShrlsService {
	var s ShrlsService

	s.SetConfig(config)

	// Set ServerStorage
	var storage *directorystate.DirectoryStorage = directorystate.New(config)
	s.SetStorage(storage)

	// Set ServerState
	state := mongostate.New(config)
	state.SetStorage(storage)
	s.SetState(state)

	// Set Server Implementation
	impl := service.New(config)
	impl.SetState(state)
	s.SetServer(impl)

	return s
}
