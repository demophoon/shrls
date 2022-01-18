package main

import (
	"fmt"
	"os"
	"time"

	"github.com/thanhpk/randstr"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ShrlType int

const (
	ShortenedUrl ShrlType = iota
	UploadedFile
	TextSnippet
)

func (t ShrlType) String() string {
	return []string{"ShortUrl", "UploadedFile", "TextSnippet"}[t]
}

type URL struct {
	ID             primitive.ObjectID `bson:"_id" json:"id"`
	Alias          string             `bson:"alias" json:"alias"`
	Location       string             `bson:"location" json:"location"`
	UploadLocation string             `bson:"upload_location" json:"-"`
	SnippetTitle   string             `bson:"snippet_title" json:"-"`
	Snippet        string             `bson:"snippet" json:"-"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	Views          int                `bson:"views" json:"views"`
	Tags           []string           `bson:"tags" json:"tags"`
	Type           ShrlType           `bson:"type" json:"type"`
}

func (u URL) Delete() error {
	switch u.Type {
	case UploadedFile:
		os.Remove(u.UploadLocation)
	}

	_, err := collection.DeleteOne(ctx, bson.M{"_id": u.ID})
	return err
}

func (u URL) IncrementViews() error {
	_, err := collection.UpdateByID(ctx, u.ID, bson.D{
		{"$inc", bson.D{{"views", 1}}},
	})
	return err

}

type URLs struct {
	Urls []*URL `json:"shrls"`
}

func NewAlias() string {
	alias := randstr.String(5)
	for aliasExists(alias) {
		alias = randstr.String(5)
	}
	return alias
}

func aliasExists(alias string) bool {
	filter := bson.D{
		primitive.E{Key: "alias", Value: alias},
	}
	urls, err := filterUrls(filter)
	if err != nil {
		return false
	}
	return len(urls) == 0
}

func NewURL() URL {
	url := URL{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		Alias:     NewAlias(),
		Type:      ShortenedUrl,
	}
	return url
}

func urlByID(url_id string) (*URL, error) {
	var url *URL
	_id, err := primitive.ObjectIDFromHex(url_id)
	if err != nil {
		return url, err
	}

	cur := collection.FindOne(ctx, bson.M{"_id": &_id})
	err = cur.Decode(&url)
	return url, err
}

func updateUrl(url *URL) error {
	_, err := collection.UpdateOne(ctx, bson.M{"_id": url.ID}, bson.D{
		{"$set", bson.D{
			{"alias", url.Alias},
			{"location", url.Location},
			{"tags", url.Tags},
		}},
	})
	return err
}

func deleteUrl(url_id string) error {
	id, err := primitive.ObjectIDFromHex(url_id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func ShrlFromString(url string) URL {
	shrl := NewURL()
	shrl.Location = url
	createUrl(&shrl)
	return shrl
}

func createUrl(url *URL) error {
	_, err := collection.InsertOne(ctx, url)
	return err
}

type PaginationParameters struct {
	Search string `json:"search"`
	Skip   int64  `json:"skip"`
	Limit  int64  `json:"limit"`
}

func paginatedUrls(prm PaginationParameters) ([]*URL, int64, error) {
	var urls []*URL

	regex := fmt.Sprintf(".*%s.*", prm.Search)

	filter := bson.D{{
		"$or",
		bson.A{
			bson.D{{
				"alias",
				bson.D{{
					"$regex",
					primitive.Regex{Pattern: regex, Options: "i"},
				}},
			}},
			bson.D{{
				"location",
				bson.D{{
					"$regex",
					primitive.Regex{Pattern: regex, Options: "i"},
				}},
			}},
			bson.D{{
				"tags",
				bson.D{{
					"$regex",
					primitive.Regex{Pattern: regex, Options: "i"},
				}},
			}},
		},
	}}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		count = -1
	}

	opts := options.FindOptions{
		Skip:  &prm.Skip,
		Limit: &prm.Limit,
	}

	opts.SetSort(bson.D{{"created_at", -1}})

	cur, err := collection.Find(ctx, filter, &opts)
	if err != nil {
		return urls, count, err
	}

	for cur.Next(ctx) {
		var u URL
		err := cur.Decode(&u)
		if err != nil {
			return urls, count, err
		}
		urls = append(urls, &u)
	}
	if err := cur.Err(); err != nil {
		return urls, count, err
	}
	return urls, count, nil
}

func filterUrls(filter interface{}) ([]*URL, error) {
	var urls []*URL

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return urls, err
	}

	for cur.Next(ctx) {
		var u URL
		err := cur.Decode(&u)
		if err != nil {
			return urls, err
		}

		urls = append(urls, &u)
	}

	if err := cur.Err(); err != nil {
		return urls, err
	}

	cur.Close(ctx)

	if len(urls) == 0 {
		return urls, mongo.ErrNoDocuments
	}

	return urls, nil
}
