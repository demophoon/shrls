package mongostate

import (
	"context"
	"os"
	"time"

	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
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

func urlToShrl(u *URL) *pb.ShortURL {
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
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_File{
				File: []byte{},
			},
		}
	case TextSnippet:
		contentType = pb.ShortURL_SNIPPET
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_Snippet{
				Snippet: &pb.Snippet{
					Title: "",
					Body:  []byte{},
				},
			},
		}
	}
	return &pb.ShortURL{
		Id:      u.ID.String(),
		Type:    contentType,
		Stub:    u.Alias,
		Content: &content,
	}
}

func shrlToUrl(u pb.ShortURL) URL {
	return URL{
		ID: [12]byte{
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
			0,
		},
		Alias:          "",
		Location:       "",
		UploadLocation: "",
		SnippetTitle:   "",
		Snippet:        "",
		CreatedAt:      time.Time{},
		Views:          0,
		Tags:           []string{},
		Type:           0,
	}
}

func (s *MongoDBState) CreateShrl(ctx context.Context, url *pb.ShortURL) (*pb.ShortURL, error) {
	//_, err := s.collection.InsertOne(ctx, u)
	return &pb.ShortURL{}, nil
}

func (s *MongoDBState) getShrl(ctx context.Context, ref *pb.Ref_ShortURL) (*URL, error) {
	var url *URL

	_id := ref.GetId()
	cur := s.collection.FindOne(ctx, bson.M{"_id": &_id})
	err := cur.Decode(&url)
	if err != nil {
		return nil, err
	}

	return url, nil
}

func (s *MongoDBState) GetShrl(ctx context.Context, ref *pb.Ref_ShortURL) (*pb.ShortURL, error) {
	url, err := s.getShrl(ctx, ref)
	if err != nil {
		return nil, err
	}
	return urlToShrl(url), nil
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
