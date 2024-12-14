package mongostate

import (
	"context"

	"github.com/demophoon/shrls/pkg/config"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBState struct {
	client     *mongo.Client
	collection *mongo.Collection
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

func New(c config.Config) *MongoDBState {
	state := &MongoDBState{}
	if c.StateBackend.Mongo == nil {
		log.Fatal("Couldn't initialize MongoDBState backend. MongoConnectionString not defined.")
	}

	if err := state.init(c.StateBackend.Mongo.ConnectionString); err != nil {
		log.Fatal("Couldn't initialize MongoDBState backend. %s", err)
	}
	return state
}

func (s MongoDBState) Close() {
}
