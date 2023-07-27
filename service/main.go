package service

import (
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
)

type Server struct {
	state   server.ServerState
	storage server.ServerStorage
}

func New(c *config.Config) *Server {
	server := &Server{}
	return server
}

func (s *Server) SetState(state server.ServerState) {
	s.state = state
}

func (s *Server) SetStorage(storage server.ServerStorage) {
	s.storage = storage
}
