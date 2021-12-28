package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goji.io/pat"
)

func shrlFromRequest(r *http.Request) *URL {
	shrl := pat.Param(r, "shrl")
	return getShrl(shrl)
}

func getShrl(shrl string) *URL {
	filter := bson.D{
		primitive.E{Key: "alias", Value: shrl},
	}
	urls, err := filterUrls(filter)
	if err != nil {
		return &URL{
			Location: "https://www.brittg.com/",
		}
	}
	return urls[rand.Intn(len(urls))]
}

func urlRedirect(w http.ResponseWriter, r *http.Request) {
	shrl := shrlFromRequest(r)
	http.Redirect(w, r, shrl.Location, 301)
}

func urlPrintInfo(w http.ResponseWriter, r *http.Request) {
	shrl := shrlFromRequest(r)
	output, err := json.Marshal(shrl)
	if err != nil {
		http.Error(w, "Invalid SHRL", 500)
	}
	w.Write(output)
}

func urlNew(w http.ResponseWriter, r *http.Request) {
	shrl := NewURL()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&shrl)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	createUrl(&shrl)
	fmt.Fprintf(w, "new url %s (%s : %s)", shrl.ID.String(), shrl.Alias, shrl.Location)
}
