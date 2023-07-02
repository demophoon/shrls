package mongostate

import (
	"context"

	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

func (s *MongoDBState) GetShrl(ctx context.Context, ref *pb.Ref_ShortURL) (pb.ShortURL, error) {
	return pb.ShortURL{
		Id:   "",
		Type: 0,
		Stub: "",
		Content: &pb.ExpandedURL{
			Content: &pb.ExpandedURL_Url{
				Url: &pb.Redirect{
					Url:     "nuts",
					Favicon: []byte{},
				},
			},
		},
	}, nil
}
