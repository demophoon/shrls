package server

import (
	pb "github.com/demophoon/shrls/server/gen"
)

type ServerHandler interface {
	pb.ShrlsServer
	pb.FileUploadServer

	SetState(ServerState)
	SetStorage(ServerStorage)
}
