package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goji.io/pat"
)

type URLUpdateResponse struct {
	Status string `json:"status"`
}

type URLListResponse struct {
	Count int64  `json:"count"`
	URLs  []*URL `json:"shrls"`
}

func shrlFromRequest(r *http.Request) (*URL, string) {
	var ext string
	shrl := pat.Param(r, "shrl")
	parts := strings.Split(shrl, ".")
	alias := parts[0]
	if len(parts) > 1 {
		ext = parts[len(parts)-1]
	}
	return getShrl(alias), ext
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

func resolveShrl(w http.ResponseWriter, r *http.Request) {
	shrl, ext := shrlFromRequest(r)
	go shrl.IncrementViews()

	switch ext {
	case "qr":
		shrl.ToQR(w)
	case "txt":
		shrl.ToText(w)
	default:
		switch shrl.Type {
		case ShortenedUrl:
			http.Redirect(w, r, shrl.Location, http.StatusPermanentRedirect)

		case UploadedFile:
			writeFile(shrl, w)

		case TextSnippet:
			w.Write([]byte(shrl.Snippet))
		}
	}
}

func writeFile(shrl *URL, w http.ResponseWriter) {
	read, err := os.Open(shrl.UploadLocation)
	if err != nil {
		http.Error(w, fmt.Sprintf("err: %v\n\n%v", err, shrl), http.StatusNotFound)
	}
	defer read.Close()

	io.Copy(w, read)
}

func urlPrintAll(w http.ResponseWriter, r *http.Request) {
	var prms PaginationParameters

	prms.Search = r.URL.Query().Get("search")
	l, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		prms.Limit = 100
	}
	s, err := strconv.ParseInt(r.URL.Query().Get("skip"), 10, 64)
	if err != nil {
		prms.Skip = 0
	}

	prms.Limit = l
	prms.Skip = s

	if prms.Limit > 100 {
		prms.Limit = 100
	}
	if prms.Limit < 25 {
		prms.Limit = 25
	}

	urls, count, err := paginatedUrls(prms)

	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to retrieve data: %s", err), 500)
	}
	pl := URLListResponse{
		Count: count,
		URLs:  urls,
	}
	//output, err := json.Marshal(pl)
	encoder := json.NewEncoder(w)
	encoder.Encode(pl)
}

func urlPrintInfo(w http.ResponseWriter, r *http.Request) {
	shrl, _ := shrlFromRequest(r)
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

	encoder := json.NewEncoder(w)
	response := URLUpdateResponse{Status: "Success"}
	encoder.Encode(response)
}

func urlModify(w http.ResponseWriter, r *http.Request) {
	shrl_id := pat.Param(r, "shrl_id")
	shrl, err := urlByID(shrl_id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to locate SHRL (%s) %s", shrl_id, err), http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var updated_shrl URL
	err = decoder.Decode(&updated_shrl)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	shrl.Alias = updated_shrl.Alias
	shrl.Location = updated_shrl.Location
	shrl.Tags = updated_shrl.Tags
	err = updateUrl(shrl)

	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	encoder := json.NewEncoder(w)
	response := URLUpdateResponse{Status: "Success"}
	encoder.Encode(response)
}

func urlDelete(w http.ResponseWriter, r *http.Request) {
	shrl_id := pat.Param(r, "shrl_id")

	response := URLUpdateResponse{Status: "Success"}
	encoder := json.NewEncoder(w)

	url, err := urlByID(shrl_id)
	if err != nil {
		response = URLUpdateResponse{Status: "Error"}
		encoder.Encode(response)
		return
	}

	err = url.Delete()
	if err != nil {
		response = URLUpdateResponse{Status: "Error"}
	}
	encoder.Encode(response)
}
