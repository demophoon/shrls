package service

import (
	"context"

	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

type Server struct{}

func (s *Server) GetShrl(ctx context.Context, req *pb.GetShrlRequest) (*pb.GetShrlResponse, error) {
	return &pb.GetShrlResponse{}, nil
}
