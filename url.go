package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL struct {
	ID        primitive.ObjectID `bson:"_id",json:"id"`
	Alias     string             `bson:"alias",json:"alias"`
	Location  string             `bson:"location",json:"location"`
	CreatedAt time.Time          `bson:"created_at",json:"created_at"`
	Views     int                `bson:"views",json:"views"`
	Tags      []*Tag             `bson:"tags",json:"tags"`
}

type Tag struct {
	ID        int       `bson:"_id",json:"id"`
	CreatedAt time.Time `bson:"created_at",json:"created_at"`
	Name      string    `bson:"name",json:"name"`
}

func NewURL() URL {
	url := URL{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
	}
	return url
}

func createUrl(url *URL) error {
	_, err := collection.InsertOne(ctx, url)
	return err
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
