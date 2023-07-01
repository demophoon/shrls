package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	shrls "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
	gw "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen/gateway"
	mongostate "gitlab.cascadia.demophoon.com/demophoon/go-shrls/state/mongo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type Service struct {
	server *shrls.ServerHandler
	state  *shrls.ServerState
}

func (s *Service) Run() error {
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

func (s Service) startHttpServer() error {
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

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Info(fmt.Sprintf("Starting HTTP Server: %d", viper.GetInt("port")))
	return http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("port")), mux)
}

func (s Service) startGRpcServer() error {
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

func (s *Service) SetState(state shrls.ServerState) {
	s.state = &state
}

func (s *Service) SetServer(server shrls.ServerHandler) {
	s.server = &server
}

// This function is responsible for returning a concrete implementation of the
// Shrls service with a mongodb state backend. Other backends can be setup by
// manually configuring the Service{} type itself.
func New() Service {
	var s Service

	// Set Server Implementation
	var impl Server
	s.SetServer(&impl)

	// Set ServerState
	mongo_uri := viper.GetString("mongo_uri")
	var mongo mongostate.MongoDBState
	err := mongo.Init(mongo_uri)
	if err != nil {
		log.Fatal("Couldn't connect to Mongo")
	}
	s.SetState(&mongo)

	return s
}
