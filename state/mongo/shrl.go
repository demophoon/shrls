package mongostate

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"time"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ShrlType int

const (
	ShortenedUrl ShrlType = iota
	UploadedFile
	TextSnippet
)

type URL struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Alias          string             `bson:"alias" json:"alias"`
	Location       string             `bson:"location" json:"location"`
	UploadLocation string             `bson:"upload_location" json:"-"`
	SnippetTitle   string             `bson:"snippet_title" json:"snippet_title,omitempty"`
	Snippet        string             `bson:"snippet" json:"snippet,omitempty"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	Views          int                `bson:"views" json:"views"`
	Tags           []string           `bson:"tags" json:"tags"`
	Type           ShrlType           `bson:"type" json:"type"`
}

func (s *MongoDBState) urlToPbShrl(u *URL) *pb.ShortURL {
	var content pb.ExpandedURL
	var contentType pb.ShortURL_ShortURLType

	switch u.Type {
	case ShortenedUrl:
		contentType = pb.ShortURL_LINK
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_Url{
				Url: &pb.Redirect{
					Url:     u.Location,
					Favicon: []byte{},
				},
			},
		}
	case UploadedFile:
		contentType = pb.ShortURL_UPLOAD
		f, err := s.storage.ReadFile(u.UploadLocation)
		if err != nil {
			log.Error("Unable to retrieve file", "URL ID", u.ID)
		}
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_File{
				File: f,
			},
		}
	case TextSnippet:
		contentType = pb.ShortURL_SNIPPET
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_Snippet{
				Snippet: &pb.Snippet{
					Title: u.SnippetTitle,
					Body:  []byte(u.Snippet),
				},
			},
		}
	}
	return &pb.ShortURL{
		Id:      u.ID.Hex(),
		Type:    contentType,
		Stub:    u.Alias,
		Content: &content,
		Tags:    u.Tags,
	}
}

func (s *MongoDBState) pbShrlToUrl(u pb.ShortURL) URL {
	url := URL{
		Alias:     u.Stub,
		Views:     int(u.Views),
		CreatedAt: time.Unix(u.CreatedAt, 0),
		Tags:      u.Tags,
	}

	id, err := primitive.ObjectIDFromHex(u.Id)
	if err != nil {
		if err != primitive.ErrInvalidHex {
			log.Error(err)
		}
	} else {
		url.ID = id
	}

	switch u.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		url.Type = ShortenedUrl
		url.Location = u.Content.GetUrl().Url
	case *pb.ExpandedURL_File:
		url.Type = UploadedFile
		key, err := s.storage.CreateFile(u.Content.GetFile())
		if err != nil {
			log.Error("Failed to create file for url", u.Id)
		} else {
			url.UploadLocation = key
		}
	case *pb.ExpandedURL_Snippet:
		url.Type = TextSnippet
		url.Snippet = string(u.Content.GetSnippet().GetBody())
		url.SnippetTitle = u.Content.GetSnippet().GetTitle()
	}

	return url
}

func (s *MongoDBState) CreateShrl(ctx context.Context, url *pb.ShortURL) (*pb.ShortURL, error) {
	u := s.pbShrlToUrl(*url)
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()
	u.Views = 0

	_, err := s.collection.InsertOne(ctx, u)
	if err != nil {
		return nil, err
	}

	return s.GetShrl(ctx, &pb.Ref_ShortURL{
		Ref: &pb.Ref_ShortURL_Id{
			Id: u.ID.Hex(),
		},
	})
}

func (s *MongoDBState) getShrl(ctx context.Context, ref *pb.Ref_ShortURL) (*URL, error) {
	urls, err := s.getShrls(ctx, ref)
	if err != nil {
		return nil, err
	}
	if len(urls) > 0 {
		return urls[rand.Intn(len(urls))], nil
	}
	return nil, fmt.Errorf("No shrls could be found")
}

func (s *MongoDBState) GetShrl(ctx context.Context, ref *pb.Ref_ShortURL) (*pb.ShortURL, error) {
	url, err := s.getShrl(ctx, ref)
	if err != nil {
		return nil, err
	}
	return s.urlToPbShrl(url), nil
}

func (s *MongoDBState) getShrls(ctx context.Context, ref *pb.Ref_ShortURL) ([]*URL, error) {
	var urls []*URL
	var query bson.M

	switch ref.Ref.(type) {
	case *gen.Ref_ShortURL_Id:
		_ids := ref.GetId()
		_id, err := primitive.ObjectIDFromHex(_ids)
		if err != nil {
			return nil, err
		}
		query = bson.M{"_id": _id}
	case *gen.Ref_ShortURL_Alias:
		query = bson.M{
			"alias": ref.GetAlias(),
		}
	}

	cur, err := s.collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var u *URL
		err := cur.Decode(&u)
		if err != nil {
			return nil, err
		}
		urls = append(urls, u)
	}

	return urls, nil
}

func (s *MongoDBState) GetShrls(ctx context.Context, ref *pb.Ref_ShortURL) ([]*pb.ShortURL, error) {
	urls, err := s.getShrls(ctx, ref)
	if err != nil {
		return nil, err
	}
	var final []*pb.ShortURL
	for _, url := range urls {
		final = append(final, s.urlToPbShrl(url))
	}
	return final, nil
}

func (s *MongoDBState) UpdateShrl(ctx context.Context, url *pb.ShortURL) (*pb.ShortURL, error) {
	//_, err := s.collection.UpdateByID(ctx, u.ID, bson.D{
	//	{"$set", u},
	//})
	//return nil, err
	return &pb.ShortURL{}, nil
}

func (s *MongoDBState) DeleteShrl(ctx context.Context, ref *pb.Ref_ShortURL) error {
	u, err := s.getShrl(ctx, ref)
	if err != nil {
		return err
	}

	switch u.Type {
	case UploadedFile:
		os.Remove(u.UploadLocation)
	}

	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": u.ID})
	return err
}
