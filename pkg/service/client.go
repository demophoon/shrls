package service

import "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"

func (s *ShrlsService) NewClient() server.ServerState {
	return s.State
}
