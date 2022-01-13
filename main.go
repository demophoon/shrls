package main

import (
	"context"
	"fmt"
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
	dbConnectionStr = "mongodb://mongo:example@localhost:27017"

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
	workdir, err := os.Getwd()
	if err != nil {
		log.Fatal(fmt.Sprintf("Error %s", err))
		os.Exit(1)
	}

	mux := goji.NewMux()

	// Admin
	fs := http.FileServer(http.Dir(workdir + "/dist/"))
	mux.Handle(pat.Get("/admin/*"), http.StripPrefix("/admin/", fs))

	// Frontend
	mux.HandleFunc(pat.Get("/:shrl"), urlRedirect)

	// Uploads
	upload_fs := http.FileServer(http.Dir(workdir + "/uploads/"))
	mux.Handle(pat.Get("/u/*"), http.StripPrefix("/u/", upload_fs))

	// Api
	mux.HandleFunc(pat.Get("/api/shrl/:shrl"), urlPrintInfo)
	mux.HandleFunc(pat.Get("/api/shrl"), urlPrintAll)
	mux.HandleFunc(pat.Put("/api/shrl/:shrl_id"), urlModify)
	mux.HandleFunc(pat.Delete("/api/shrl/:shrl_id"), urlDelete)
	mux.HandleFunc(pat.Post("/api/shrl"), urlNew)

	http.ListenAndServe(":8000", mux)
}
