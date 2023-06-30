package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	shrls "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
	mongostate "gitlab.cascadia.demophoon.com/demophoon/go-shrls/state/mongo"
)

type Service struct {
	server *shrls.ServerHandler
	state  *shrls.ServerState
}

func (s *Service) Run() error {
	log.Info("Server running: ", viper.Get("port"))
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

	// Set ServerState
	mongo_uri := viper.GetString("mongo_uri")
	var mongo *mongostate.MongoDBState
	mongo.Init(mongo_uri)
	s.SetState(mongo)

	return s
}
