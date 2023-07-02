package server

import (
	"context"

	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

type ServerState interface {
	GetShrl(context.Context, *pb.Ref_ShortURL) (pb.ShortURL, error)
}
