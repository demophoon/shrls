package serve

import (
	"context"

	"github.com/demophoon/shrls/pkg/config"
	"github.com/demophoon/shrls/pkg/service"
	pb "github.com/demophoon/shrls/server/gen"
	directorystate "github.com/demophoon/shrls/storage/directory"

	log "github.com/sirupsen/logrus"
)

func setupDemo() {
	// Seed data
	ctx := context.Background()
	defer ctx.Done()

	config := config.New()
	s := service.New(config)
	d := directorystate.New(config)

	client := s.NewClient()
	defer client.Close()
	client.CreateShrl(ctx, &pb.ShortURL{
		Content: &pb.ExpandedURL{
			Content: &pb.ExpandedURL_Url{
				Url: &pb.Redirect{
					Url: "https://www.brittg.com/",
				},
			},
		},
	})
	client.CreateShrl(ctx, &pb.ShortURL{
		Stub: "resume",
		Content: &pb.ExpandedURL{
			Content: &pb.ExpandedURL_Url{
				Url: &pb.Redirect{
					Url: "https://resume.brittg.com/",
				},
			},
		},
	})
	client.CreateShrl(ctx, &pb.ShortURL{
		Content: &pb.ExpandedURL{
			Content: &pb.ExpandedURL_Snippet{
				Snippet: &pb.Snippet{
					Title: "Example Snippet",
					Body:  []byte("This is an example snippet!"),
				},
			},
		},
	})
	ref, size, err := d.CreateFile([]byte("This is an example of an uploaded file"))
	if err != nil {
		log.Errorf("Unable to upload example file: %s", err)
	}
	client.CreateShrl(ctx, &pb.ShortURL{
		Content: &pb.ExpandedURL{
			Content: &pb.ExpandedURL_File{
				File: &pb.Upload{
					Ref: ref,
					Size: size,
				},
			},
		},
	})
}
