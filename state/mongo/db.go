package mongostate

import (
	"context"
	"fmt"
	"log"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"
	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/server"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBState struct {
	client     *mongo.Client
	collection *mongo.Collection
	storage    server.ServerStorage
}

func (s *MongoDBState) init(conn string) error {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(conn)
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	s.client = c
	s.collection = s.client.Database("shrls").Collection("urls")

	return nil
}

func New(c *config.Config) *MongoDBState {
	log.Printf(fmt.Sprintf("Config db: %v", c))
	state := &MongoDBState{}
	if err := state.init(c.MongoConnectionString); err != nil {
		log.Fatal("Couldn't initialize MongoDBState. %s", err)
	}
	return state
}

func (s *MongoDBState) SetStorage(storage server.ServerStorage) {
	s.storage = storage
}
