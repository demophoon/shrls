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
	"gopkg.in/yaml.v3"

	"github.com/goji/httpauth"
	"goji.io"
)

type ShrlSettings struct {
	BaseURL               string
	Port                  int
	DefaultRedirect       string
	UploadDirectory       string
	MongoUsername         string
	MongoPassword         string
	MongoConnectionString string
	AdminUsername         string
	AdminPassword         string

	SettingsFilepath       string
	ResolveLocationHosts   []string `yaml:"resolveLocation"`
	StripQueryParamsHosts  []string `yaml:"stripQueryParams"`
	TerminalRedirect       bool     `yaml:"terminalRedirect"`
	TerminalRedirectString string   `yaml:"terminalRedirectString"`
}

func (s *ShrlSettings) Parse(data []byte) error {
	return yaml.Unmarshal(data, s)
}

var collection *mongo.Collection
var ctx = context.TODO()

var Settings ShrlSettings

func init() {
	// Init Settings
	port := 8000

	Settings = ShrlSettings{
		BaseURL:         os.Getenv("SHRLS_BASE_URL"),
		Port:            port,
		DefaultRedirect: os.Getenv("DEFAULT_REDIRECT"),
		UploadDirectory: "/local",
		// mongodb://mongo:example@localhost:27017
		MongoUsername:         os.Getenv("MONGO_USERNAME"),
		MongoPassword:         os.Getenv("MONGO_PASSWORD"),
		MongoConnectionString: os.Getenv("MONGO_URI"),
		AdminUsername:         os.Getenv("SHRLS_USERNAME"),
		AdminPassword:         os.Getenv("SHRLS_PASSWORD"),
		SettingsFilepath:      os.Getenv("SHRLS_SETTINGS_FILE"),
		TerminalRedirect:      false,
	}

	//Settings.MongoConnectionString = fmt.Sprintf("mongodb://%s:%s@10.211.55.6/shrls", Settings.MongoUsername, Settings.MongoPassword)

	fmt.Printf("settings: %v#\n", Settings)

	// Init Mongo
	clientOptions := options.Client().ApplyURI(Settings.MongoConnectionString)
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
	admin_mux := goji.SubMux()
	api_mux := goji.SubMux()
	curl_api_mux := goji.SubMux()

	if Settings.AdminUsername != "" && Settings.AdminPassword != "" {
		authOpts := httpauth.AuthOptions{
			User:                Settings.AdminUsername,
			Password:            Settings.AdminPassword,
			UnauthorizedHandler: http.HandlerFunc(UnauthorizedWarning),
		}
		auth_middleware := httpauth.BasicAuth(authOpts)
		api_mux.Use(auth_middleware)
		admin_mux.Use(auth_middleware)
		curl_api_mux.Use(auth_middleware)
	}

	// Admin
	mux.Handle(pat.New("/admin/*"), admin_mux)
	fs := http.FileServer(http.Dir(workdir + "/dist/"))
	admin_mux.Handle(pat.Get("/*"), http.StripPrefix("/admin/", fs))

	// Api
	mux.Handle(pat.New("/api/*"), api_mux)
	api_mux.HandleFunc(pat.Get("/shrl/:shrl"), urlPrintInfo)
	api_mux.HandleFunc(pat.Get("/shrl"), urlPrintAll)
	api_mux.HandleFunc(pat.Put("/shrl/:shrl_id"), urlModify)
	api_mux.HandleFunc(pat.Delete("/shrl/:shrl_id"), urlDelete)
	api_mux.HandleFunc(pat.Post("/shrl"), urlNew)

	// Bookmarklet API
	api_mux.HandleFunc(pat.Get("/bookmarklet/new"), bookmarkletNew)

	// File Uploads
	api_mux.HandleFunc(pat.Post("/upload"), fileUpload)

	// Snippets
	api_mux.HandleFunc(pat.Post("/snippet"), snippetUpload)
	api_mux.HandleFunc(pat.Get("/snippet/:snippet_id"), snippetGet)

	// Frontend
	mux.HandleFunc(pat.Get("/"), defaultRedirect)
	mux.HandleFunc(pat.Get("/:shrl"), resolveShrl)
	mux.HandleFunc(pat.Get("/:shrl/:search"), resolveShrl)

	// Curl API
	mux.Handle(pat.New("/*"), curl_api_mux)
	curl_api_mux.HandleFunc(pat.Post("/"), curlNew)

	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", Settings.Port), mux)
}
