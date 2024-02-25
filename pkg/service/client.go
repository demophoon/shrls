package service

import "github.com/demophoon/shrls/server"

func (s *ShrlsService) NewClient() server.ServerState {
	return s.State
}
