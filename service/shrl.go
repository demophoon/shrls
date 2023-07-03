package service

import (
	"context"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

type Server struct {
	state *server.ServerState
}

func New() *Server {
	return &Server{}
}

func (s *Server) SetState(state server.ServerState) {
	s.state = &state
}

func (s Server) GetShrl(ctx context.Context, req *pb.GetShrlRequest) (*pb.GetShrlResponse, error) {
	state := *s.state
	shrl, err := state.GetShrl(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}

	return &pb.GetShrlResponse{
		Shrl: shrl,
	}, nil
}
