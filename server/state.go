package server

import (
	"context"

	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

type ServerState interface {
	SetStorage(ServerStorage)

	GetShrl(context.Context, *pb.Ref_ShortURL) (*pb.ShortURL, error)
	GetShrls(context.Context, *pb.Ref_ShortURL) ([]*pb.ShortURL, error)
	CreateShrl(context.Context, *pb.ShortURL) (*pb.ShortURL, error)
}
