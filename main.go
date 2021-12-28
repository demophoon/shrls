package main

import (
	"context"
	"net/http"
	"os"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"goji.io/pat"

	"goji.io"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	// mongodb://mongo:example@localhost:27017
	dbConnectionStr := os.Getenv("MONGO_URI")

	// Init Mongo
	clientOptions := options.Client().ApplyURI(dbConnectionStr)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("shrls").Collection("urls")
}

func main() {
	mux := goji.NewMux()

	fs := http.FileServer(http.Dir("/mnt/c/Users/demop/projects/go-shrls/dist/"))

	// Admin
	mux.Handle(pat.Get("/admin/*"), http.StripPrefix("/admin/", fs))

	// Frontend
	mux.HandleFunc(pat.Get("/:shrl"), urlRedirect)

	// Api
	mux.HandleFunc(pat.Get("/api/shrl/:shrl"), urlPrintInfo)
	mux.HandleFunc(pat.Post("/api/shrl"), urlNew)

	http.ListenAndServe("localhost:8000", mux)
}
