package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"goji.io/pat"
)

type SnippetRequest struct {
	SnippetTitle string `json:"title"`
	SnippetBody  string `json:"body"`
}

type SnippetInformationResponse struct {
	ID    primitive.ObjectID `json:"id"`
	Title string             `json:"title"`
	Body  string             `json:"body"`
}

func uploadSnippet(snippet SnippetRequest) URL {
	shrl := NewURL()
	shrl.SnippetTitle = snippet.SnippetTitle
	shrl.Snippet = snippet.SnippetBody
	shrl.Type = TextSnippet
	shrl.Create()
	return shrl
}

func snippetUpload(w http.ResponseWriter, r *http.Request) {
	var snippet SnippetRequest
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	decoder.Decode(&snippet)
	url := uploadSnippet(snippet)

	SuccessResponse(w, &url)
}

func snippetGet(w http.ResponseWriter, r *http.Request) {
	snippet_id := pat.Param(r, "snippet_id")
	shrl, err := urlByID(snippet_id)
	if err != nil {
		http.Error(w, "Invalid URL", http.StatusNotFound)
	}

	var snippet SnippetInformationResponse
	snippet.ID = shrl.ID
	snippet.Title = shrl.SnippetTitle
	snippet.Body = shrl.Snippet

	encoder := json.NewEncoder(w)
	encoder.Encode(snippet)
}
