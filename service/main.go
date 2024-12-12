package service

import (
	boltstate "github.com/demophoon/shrls/state/boltdb"
	mongostate "github.com/demophoon/shrls/state/mongo"

	"github.com/demophoon/shrls/pkg/config"
	"github.com/demophoon/shrls/server"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	state   server.ServerState
	storage server.ServerStorage
}

func New(c *config.Config) *Server {
	server := &Server{}
	server.setState(c)
	return server
}

func (s *Server) setState(c *config.Config) {
	if c.StateBackend == nil {
		log.Fatal("State is not configured")
	}

	var cfgState server.ServerState
	if c.StateBackend.Mongo != nil {
		log.Debug("Setting up mongodb state")
		cfgState = mongostate.New(*c)
	}
	if c.StateBackend.Bolt != nil {
		log.Debug("Setting up boltdb state")
		cfgState = boltstate.New(*c)
	}

	s.state = cfgState
}

func (s *Server) setStorage(storage server.ServerStorage) {
	s.storage = storage
}

func (s *Server) SetState(state server.ServerState) {
	s.state = state
}

func (s *Server) SetStorage(storage server.ServerStorage) {
	s.storage = storage
}
