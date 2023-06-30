package server

import (
	"context"

	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

type ServerHandler interface {
	GetShrl(context.Context, pb.GetShrlRequest) (pb.GetShrlResponse, error)
}
