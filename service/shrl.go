package service

import (
	"context"

	pb "github.com/demophoon/shrls/server/gen"
	"github.com/mat/besticon/v3/besticon"
	log "github.com/sirupsen/logrus"
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

func (s Server) resolveFavicon(u *pb.ShortURL) {
	switch u.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		loc := u.Content.GetUrl().Url

		bi := besticon.New()
		icf := bi.NewIconFinder()

		icons, err := icf.FetchIcons(loc)

		if err == nil {
			for _, icon := range icons {
				if icon.Width >= 16 && icon.Height >= 16 {
					u.Content.Content = &pb.ExpandedURL_Url{
						Url: &pb.Redirect{
							Url:     loc,
							Favicon: icon.ImageData,
						},
					}
					log.Tracef("favicon found for %s at %s", loc, icon.URL)
					break
				}
			}
		} else {
			log.Warnf("Unable to fetch favicon from %s: %s", loc, err)
		}
	}
}

func (s Server) PostShrl(ctx context.Context, req *pb.PostShrlRequest) (*pb.PostShrlResponse, error) {
	u := req.Shrl
	s.resolveFavicon(u)

	shrl, err := s.state.CreateShrl(ctx, u)
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
