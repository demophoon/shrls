package service

import (
	"context"

	pb "github.com/demophoon/shrls/server/gen"
)

func (s Server) PostFileUpload(ctx context.Context, req *pb.PostFileUploadRequest) (*pb.PostFileUploadResponse, error) {
	file := req.File
	key, _, err := s.storage.CreateFile(file)
	if err != nil {
		return nil, err
	}
	return &pb.PostFileUploadResponse{
		File: &pb.Ref_FileUpload{
			Ref: &pb.Ref_FileUpload_Id{
				Id: key,
			},
		},
	}, nil
}
