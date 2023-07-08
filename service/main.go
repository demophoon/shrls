package service

import (
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
)

type Server struct {
	state *server.ServerState
}

func New(c *config.Config) *Server {
	return &Server{}
}

func (s *Server) SetState(state server.ServerState) {
	s.state = &state
}
