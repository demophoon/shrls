package service

import (
	"context"

	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

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

func (s Server) GetShrls(ctx context.Context, req *pb.GetShrlsRequest) (*pb.GetShrlsResponse, error) {
	state := *s.state
	shrl, err := state.GetShrls(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}

	return &pb.GetShrlsResponse{
		Shrl: shrl,
	}, nil
}
