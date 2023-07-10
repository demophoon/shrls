package service

import (
	"context"
	"fmt"
	"io/fs"
	"net"
	"net/http"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
	shrls "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
	gw "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen/gateway"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/service"
	mongostate "gitlab.cascadia.demophoon.com/demophoon/go-shrls/state/mongo"
	directorystate "gitlab.cascadia.demophoon.com/demophoon/go-shrls/storage/directory"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/ui"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ShrlsService struct {
	server  *shrls.ServerHandler
	State   *shrls.ServerState
	storage *shrls.ServerStorage
	config  *config.Config
}

func (s *ShrlsService) Run() error {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pb.RegisterShrlsServer(grpcServer, *s.server)

	mux := runtime.NewServeMux()
	err := gw.RegisterShrlsHandlerFromEndpoint(context.Background(), mux, fmt.Sprintf("localhost:%d", s.config.Port), []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		return err
	}

	sub, err := fs.Sub(ui.Content, "dist")
	if err != nil {
		panic(err)
	}
	fs := http.FileServer(http.FS(sub))

	main := http.NewServeMux()
	main.Handle("/admin/", http.StripPrefix("/admin/", fs))
	main.Handle("/v1/", mux)
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

	go server.Serve(httpL)
	go grpcServer.Serve(grpcL)

	return m.Serve()
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
