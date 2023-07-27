package server

import (
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

type ServerHandler interface {
	pb.ShrlsServer
	pb.FileUploadServer

	SetState(ServerState)
	SetStorage(ServerStorage)
}
