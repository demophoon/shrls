package service

import (
	"context"

	pb "github.com/demophoon/shrls/server/gen"
)

func (s Server) GetShrl(ctx context.Context, req *pb.GetShrlRequest) (*pb.GetShrlResponse, error) {
	shrl, err := s.state.GetShrl(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}

	return &pb.GetShrlResponse{
		Shrl: shrl,
	}, nil
}

func (s Server) GetShrls(ctx context.Context, req *pb.GetShrlsRequest) (*pb.GetShrlsResponse, error) {
	shrls, err := s.state.GetShrls(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}

	return &pb.GetShrlsResponse{
		Shrls: shrls,
	}, nil
}

func (s Server) ListShrls(ctx context.Context, req *pb.ListShrlsRequest) (*pb.ListShrlsResponse, error) {
	shrls, total, err := s.state.ListShrls(ctx, req.Search, req.Count, req.Page)
	if err != nil {
		return nil, err
	}

	return &pb.ListShrlsResponse{
		Shrls:      shrls,
		TotalShrls: total,
	}, nil
}

func (s Server) PutShrl(ctx context.Context, req *pb.PutShrlRequest) (*pb.PutShrlResponse, error) {
	shrl, err := s.state.UpdateShrl(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}

	return &pb.PutShrlResponse{
		Shrl: shrl,
	}, nil
}

func (s Server) PostShrl(ctx context.Context, req *pb.PostShrlRequest) (*pb.PostShrlResponse, error) {
	shrl, err := s.state.CreateShrl(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}
	return &pb.PostShrlResponse{
		Shrl: shrl,
	}, nil
}

func (s Server) DeleteShrl(ctx context.Context, req *pb.DeleteShrlRequest) (*pb.DeleteShrlResponse, error) {
	shrl, err := s.state.GetShrl(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}

	// Delete files uploaded via the FileUpload service
	switch shrl.Content.Content.(type) {
	case *pb.ExpandedURL_File:
		s.storage.DeleteFile(shrl.Content.GetFile().Ref)
	}

	err = s.state.DeleteShrl(ctx, req.Shrl)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteShrlResponse{}, nil
}
