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
	URL    *URL   `json:"shrl"`
}

type URLListResponse struct {
	Count int64  `json:"count"`
	URLs  []*URL `json:"shrls"`
}

func ErrorResponse(w http.ResponseWriter, url *URL, status int) {
	StatusResponse(w, url, "Error", status)
}

func SuccessResponse(w http.ResponseWriter, url *URL) {
	StatusResponse(w, url, "Success", http.StatusOK)
}

func StatusResponse(w http.ResponseWriter, url *URL, status string, code int) {
	if url == nil {
		url = &URL{}
	}
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	response := URLUpdateResponse{Status: status, URL: url}
	encoder.Encode(response)
}

func defaultRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, Settings.DefaultRedirect, http.StatusTemporaryRedirect)
}

func shrlFromRequest(r *http.Request) (*URL, string, error) {
	var ext string
	shrl := pat.Param(r, "shrl")
	parts := strings.Split(shrl, ".")
	alias := parts[0]
	if len(parts) > 1 {
		ext = parts[len(parts)-1]
	}
	url, err := getShrl(alias)
	if err != nil {
		return url, ext, err
	}
	return url, ext, err
}

func getShrl(shrl string) (*URL, error) {
	var url *URL
	filter := bson.D{
		primitive.E{Key: "alias", Value: shrl},
	}
	urls, err := filterUrls(filter)
	if err != nil {
		return url, err
	}
	return urls[rand.Intn(len(urls))], nil
}

func resolveShrl(w http.ResponseWriter, r *http.Request) {
	shrl, ext, err := shrlFromRequest(r)
	if err != nil {
		defaultRedirect(w, r)
		return
	}
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
	shrl, _, err := shrlFromRequest(r)
	if err != nil {
		ErrorResponse(w, shrl, http.StatusInternalServerError)
	}
	output, err := json.Marshal(shrl)
	if err != nil {
		ErrorResponse(w, shrl, http.StatusInternalServerError)
	}
	w.Write(output)
}

func urlNew(w http.ResponseWriter, r *http.Request) {
	shrl := NewURL()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&shrl)
	if err != nil {
		ErrorResponse(w, nil, http.StatusBadRequest)
		return
	}
	createUrl(&shrl)

	SuccessResponse(w, &shrl)
}

func urlModify(w http.ResponseWriter, r *http.Request) {
	shrl_id := pat.Param(r, "shrl_id")
	shrl, err := urlByID(shrl_id)
	if err != nil {
		ErrorResponse(w, nil, http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var updated_shrl URL
	err = decoder.Decode(&updated_shrl)
	if err != nil {
		ErrorResponse(w, shrl, http.StatusBadRequest)
		return
	}

	shrl.Alias = updated_shrl.Alias
	shrl.Tags = updated_shrl.Tags

	switch shrl.Type {
	case ShortenedUrl:
		shrl.Location = updated_shrl.Location
	case TextSnippet:
		shrl.Snippet = updated_shrl.Snippet
		shrl.SnippetTitle = updated_shrl.SnippetTitle
	}
	err = updateUrl(shrl)

	if err != nil {
		ErrorResponse(w, shrl, http.StatusBadRequest)
		return
	}
	SuccessResponse(w, shrl)
}

func urlDelete(w http.ResponseWriter, r *http.Request) {
	shrl_id := pat.Param(r, "shrl_id")

	url, err := urlByID(shrl_id)
	if err != nil {
		ErrorResponse(w, nil, http.StatusNotFound)
		return
	}

	err = url.Delete()
	if err != nil {
		ErrorResponse(w, url, http.StatusInternalServerError)
		return
	}
	SuccessResponse(w, url)
}
