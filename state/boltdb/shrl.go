package boltstate

import (
	"context"
	"fmt"
	"time"

	pb "github.com/demophoon/shrls/server/gen"

	storm "github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
)

type ShrlType int

const (
	ShortenedUrl ShrlType = iota
	UploadedFile
	TextSnippet
)

type URL struct {
	ID             uuid.UUID `storm:"index"`
	Alias          string
	Location       string
	UploadLocation string
	SnippetTitle   string
	Snippet        string
	CreatedAt      time.Time `storm:"index"`
	Views          int
	Tags           []string
	Type           ShrlType
}

func (s *BoltDBState) urlsToPbShrls(urls []*URL) []*pb.ShortURL {
	var final []*pb.ShortURL
	for _, url := range urls {
		final = append(final, s.urlToPbShrl(url))
	}
	return final
}

func (s *BoltDBState) pbShrlsToUrls(urls []*pb.ShortURL) []*URL {
	var final []*URL
	for _, url := range urls {
		final = append(final, s.pbShrlToUrl(url))
	}
	return final
}

func (s *BoltDBState) pbShrlToUrl(in *pb.ShortURL) *URL {
	out := URL{}
	uuid, _ := uuid.Parse(in.Id)
	out.ID = uuid
	out.Views = int(in.Views)
	out.Tags = in.Tags
	out.Alias = in.Stub

	switch in.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		out.Type = ShortenedUrl
		out.Location = in.Content.GetUrl().GetUrl()
	case *pb.ExpandedURL_File:
		out.Type = UploadedFile
		out.UploadLocation = in.Content.GetFile().GetRef()
	case *pb.ExpandedURL_Snippet:
		out.Type = TextSnippet
		out.Snippet = string(in.Content.GetSnippet().GetBody())
		out.SnippetTitle = in.Content.GetSnippet().GetTitle()
	default:
		panic("Invalid url type")
	}

	return &out
}

func (s *BoltDBState) urlToPbShrl(in *URL) *pb.ShortURL {
	var content pb.ExpandedURL
	var contentType pb.ShortURL_ShortURLType

	switch in.Type {
	case ShortenedUrl:
		contentType = pb.ShortURL_LINK
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_Url{
				Url: &pb.Redirect{
					Url:     in.Location,
					Favicon: []byte{},
				},
			},
		}
	case UploadedFile:
		contentType = pb.ShortURL_UPLOAD
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_File{
				File: &pb.Upload{
					Ref: in.UploadLocation,
				},
			},
		}
	case TextSnippet:
		contentType = pb.ShortURL_SNIPPET
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_Snippet{
				Snippet: &pb.Snippet{
					Title: in.SnippetTitle,
					Body:  []byte(in.Snippet),
				},
			},
		}
	}
	return &pb.ShortURL{
		Id:        in.ID.String(),
		Type:      contentType,
		Stub:      in.Alias,
		Content:   &content,
		Tags:      in.Tags,
		Views:     int64(in.Views),
		CreatedAt: in.CreatedAt.Unix(),
	}
}

func (s *BoltDBState) getShrl(ctx context.Context, ref *pb.Ref_ShortURL) (*URL, error) {

	var url URL

	switch ref.Ref.(type) {
	case *pb.Ref_ShortURL_Id:
		uid, err := uuid.Parse(ref.GetId())
		if err != nil {
			return nil, err
		}
		err = s.db.One("ID", uid, &url)
		if err != nil {
			return nil, err
		}
	case *pb.Ref_ShortURL_Alias:
		err := s.db.One("Alias", ref.GetAlias(), &url)
		if err != nil {
			return nil, err
		}
	}

	return &url, nil
}

func (s *BoltDBState) GetShrl(ctx context.Context, ref *pb.Ref_ShortURL) (*pb.ShortURL, error) {
	url, err := s.getShrl(ctx, ref)
	if err != nil {
		return nil, err
	}

	return s.urlToPbShrl(url), nil
}

func (s *BoltDBState) GetShrls(ctx context.Context, ref *pb.Ref_ShortURL) ([]*pb.ShortURL, error) {
	var urls []*URL
	var query storm.Query

	switch ref.Ref.(type) {
	case *pb.Ref_ShortURL_Id:
		uid, err := uuid.Parse(ref.GetId())
		if err != nil {
			return nil, err
		}
		query = s.db.Select(q.Eq("ID", uid))
	case *pb.Ref_ShortURL_Alias:
		query = s.db.Select(q.Eq("Alias", ref.GetAlias()))
	}

	err := query.Find(&urls)
	if err != nil {
		return nil, err
	}

	return s.urlsToPbShrls(urls), nil
}

func (s *BoltDBState) newStub() string {
	return randstr.String(5)
}

func (s *BoltDBState) CreateShrl(ctx context.Context, shrl *pb.ShortURL) (*pb.ShortURL, error) {
	u := s.pbShrlToUrl(shrl)
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.Views = 0
	u.Alias = s.newStub()

	err := s.db.Save(u)
	if err != nil {
		return nil, err
	}
	out := s.urlToPbShrl(u)
	return out, nil
}

func (s *BoltDBState) ListShrls(ctx context.Context, search *string, count *int64, page *int64) ([]*pb.ShortURL, int64, error) {
	var final []*pb.ShortURL
	var urls []*URL
	err := s.db.AllByIndex("CreatedAt", &urls)
	if err != nil {
		return final, -1, err
	}
	final = s.urlsToPbShrls(urls)
	return final, -1, nil
}

func (s *BoltDBState) UpdateShrl(ctx context.Context, shrl *pb.ShortURL) (*pb.ShortURL, error) {
	url := s.pbShrlToUrl(shrl)

	log.Trace(fmt.Sprintf("url:%#v", url))

	err := s.db.Update(url)
	if err != nil {
		return nil, err
	}
	return s.urlToPbShrl(url), nil
}

func (s *BoltDBState) DeleteShrl(ctx context.Context, ref *pb.Ref_ShortURL) error {
	url, err := s.getShrl(ctx, ref)
	if err != nil {
		return err
	}
	return s.db.DeleteStruct(url)
}
