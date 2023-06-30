package mongostate

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBState struct {
	client *mongo.Client
}

func (s *MongoDBState) Init(conn string) error {
	ctx := context.TODO()

	options.Client().ApplyURI(conn)
	client, err := mongo.Connect(ctx)
	if err != nil {
		return err
	}
	s.client = client
	return nil
}
