package mongostate

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"

	log "github.com/sirupsen/logrus"
	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShrlType int

const (
	ShortenedUrl ShrlType = iota
	UploadedFile
	TextSnippet
)

var (
	DefaultSearch string = ".*"
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

func (s *MongoDBState) urlsToPbShrls(urls []*URL) []*pb.ShortURL {
	var final []*pb.ShortURL
	for _, url := range urls {
		final = append(final, s.urlToPbShrl(url))
	}
	return final
}

func (s *MongoDBState) pbShrlsToUrls(urls []*pb.ShortURL) []*URL {
	var final []*URL
	for _, url := range urls {
		final = append(final, s.pbShrlToUrl(url))
	}
	return final
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
		content = pb.ExpandedURL{
			Content: &pb.ExpandedURL_File{
				File: &pb.Upload{
					Ref: u.UploadLocation,
				},
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
		Id:        u.ID.Hex(),
		Type:      contentType,
		Stub:      u.Alias,
		Content:   &content,
		Tags:      u.Tags,
		Views:     int64(u.Views),
		CreatedAt: u.CreatedAt.Unix(),
	}
}

func (s *MongoDBState) pbShrlToUrl(u *pb.ShortURL) *URL {
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
		url.UploadLocation = u.Content.GetFile().GetRef()
	case *pb.ExpandedURL_Snippet:
		url.Type = TextSnippet
		url.Snippet = string(u.Content.GetSnippet().GetBody())
		url.SnippetTitle = u.Content.GetSnippet().GetTitle()
	default:
		panic("Invalid url type")
	}

	return &url
}

func (s *MongoDBState) newStub() string {
	return randstr.String(5)
}

func (s *MongoDBState) CreateShrl(ctx context.Context, url *pb.ShortURL) (*pb.ShortURL, error) {
	u := s.pbShrlToUrl(url)
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()
	u.Views = 0
	u.Alias = s.newStub()

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
	return s.urlsToPbShrls(urls), nil
}

func (s *MongoDBState) updateShrl(ctx context.Context, url *URL) (*URL, error) {
	_, err := s.collection.UpdateByID(ctx, url.ID, bson.M{
		"$set": url,
	})
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (s *MongoDBState) UpdateShrl(ctx context.Context, url *pb.ShortURL) (*pb.ShortURL, error) {
	u, err := s.updateShrl(ctx, s.pbShrlToUrl(url))
	if err != nil {
		return nil, err
	}
	return s.urlToPbShrl(u), nil
}

func (s *MongoDBState) listShrls(ctx context.Context, search *string, count *int64, page *int64) ([]*URL, int64, error) {
	filter := bson.M{}
	if search == nil {
		search = &DefaultSearch
	}

	pregex := bson.M{
		"$regex": primitive.Regex{
			Pattern: *search,
			Options: "i",
		},
	}

	filter = bson.M{
		"$or": bson.A{
			bson.M{"alias": pregex},
			bson.M{"location": pregex},
			bson.M{"tags": pregex},
		},
	}

	total, err := s.collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, -1, err
	}

	var searchCount int64 = 50
	var searchPage int64 = 0

	if count != nil {
		searchCount = *count
	}
	if searchCount > 100 {
		searchCount = 100
	}
	if searchCount < 10 {
		searchCount = 10
	}

	if page != nil {
		searchPage = *page
	}
	if searchPage < 0 {
		searchPage = 0
	}

	skip := searchCount * searchPage
	limit := searchCount

	opts := options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	opts.SetSort(bson.M{"created_at": -1})

	cur, err := s.collection.Find(ctx, filter, &opts)

	urls := []*URL{}

	for cur.Next(ctx) {
		var u URL
		err := cur.Decode(&u)
		if err != nil {
			log.Error("Couldn't decode SHRL from Mongo", err)
			continue
		}
		urls = append(urls, &u)
	}

	return urls, total, nil
}

func (s *MongoDBState) ListShrls(ctx context.Context, search *string, count *int64, page *int64) ([]*pb.ShortURL, int64, error) {
	urls, total, err := s.listShrls(ctx, search, count, page)

	return s.urlsToPbShrls(urls), total, err
}

func (s *MongoDBState) DeleteShrl(ctx context.Context, ref *pb.Ref_ShortURL) error {
	u, err := s.getShrl(ctx, ref)
	if err != nil {
		return err
	}

	_, err = s.collection.DeleteOne(ctx, bson.M{"_id": u.ID})
	return err
}
