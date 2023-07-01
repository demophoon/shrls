package service

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/spf13/viper"

	shrls "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
	gw "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen/gateway"
	mongostate "gitlab.cascadia.demophoon.com/demophoon/go-shrls/state/mongo"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Service struct {
	server *shrls.ServerHandler
	state  *shrls.ServerState
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func (s *Service) Run() error {
	//log.Info("Server running: ", viper.Get("port"))

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := gw.RegisterShrlsHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
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
