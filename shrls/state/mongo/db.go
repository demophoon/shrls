package mongostate

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBState struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func (s MongoDBState) Init(conn string) error {
	ctx := context.TODO()
	options.Client().ApplyURI(conn)
	c, err := mongo.Connect(ctx)
	if err != nil {
		return err
	}
	s.client = c
	s.collection = s.client.Database("shrls").Collection("urls")

	return nil
}
