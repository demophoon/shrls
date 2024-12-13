package service

import (
	"github.com/demophoon/shrls/pkg/config"
	"github.com/demophoon/shrls/server"
)

type Server struct {
	state   server.ServerState
	storage server.ServerStorage
	config  *config.Config
}

func New(c *config.Config) *Server {
	server := &Server{
		config: c,
	}
	return server
}

func (s *Server) SetState(state server.ServerState) {
	s.state = state
}

func (s *Server) SetStorage(storage server.ServerStorage) {
	s.storage = storage
}
