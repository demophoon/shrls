package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
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

func isUA(r *http.Request, agent string) bool {
	ua := strings.ToLower(r.UserAgent())
	return strings.HasPrefix(ua, strings.ToLower(agent))
}

func isTerminal(r *http.Request) bool {
	return isUA(r, "curl") || isUA(r, "wget")
}

func defaultRedirect(w http.ResponseWriter, r *http.Request) {
	if isTerminal(r) && Settings.TerminalRedirect && Settings.TerminalRedirectString != "" {
		w.Write([]byte(Settings.TerminalRedirectString + "\n"))
		return
	}
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
	if shrl.Alias == ".well-known" {
		webfinger(w, r)
		return
	}
	go shrl.IncrementViews()

	switch ext {
	case "qr", "qrcode":
		if isTerminal(r) {
			shrl.toTextQR(w)
		} else {
			shrl.ToQR(w)
		}
	case "txt", "text":
		shrl.ToText(w)
	default:
		shrl.Redirect(w, r)
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

func bookmarkletNew(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("u")
	img := r.URL.Query().Get("i")
	snippet := r.URL.Query().Get("s")
	var shrl URL
	if url != "" {
		shrl = ShrlFromString(url)
	} else if img != "" {
		by, err := base64.StdEncoding.DecodeString(img)
		if err != nil {
			return
		}
		file_uuid, _ := uuid.New()
		filepath := path.Join(Settings.UploadDirectory, fmt.Sprintf("%s", file_uuid))
		dst, _ := os.Create(filepath)
		defer dst.Close()
		dst.Write(by)
		shrl = FileFromString(filepath)
	} else if snippet != "" {
		shrl = uploadSnippet(SnippetRequest{
			SnippetTitle: "From Bookmarklet",
			SnippetBody:  snippet,
		})
	}

	shrljson, _ := json.Marshal(&struct {
		Info URL    `json:"info"`
		Url  string `json:"shrl"`
	}{
		Info: shrl,
		Url:  shrl.FriendlyAlias(),
	})
	js_preamble := fmt.Sprintf(`
	document.currentScript.dispatchEvent(new CustomEvent('shrls-response', {
		detail: %s,
	}))
	`, shrljson)
	w.Header().Set("Content-Type", "application/javascript")
	w.Write([]byte(js_preamble))
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

func curlNew(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	mtype := mimetype.Detect(b)

	var shrl URL
	if mtype.Is("text/plain") {
		body := string(b)
		if strings.HasPrefix(strings.ToLower(body), "http") && !strings.ContainsAny(body, "\n") {
			shrl = ShrlFromString(body)
		} else {
			shrl = uploadSnippet(SnippetRequest{
				SnippetBody: body,
			})
		}
	} else {
		s, err := URLFromFile(b)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error"))
			return
		}
		shrl = s
	}

	baseUrl := r.Host + "/"

	w.Write([]byte(baseUrl + shrl.Alias + "\n"))
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
	shrl.Create()

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

func UnauthorizedWarning(w http.ResponseWriter, r *http.Request) {

	host := r.Host

	unauthorized_message := "Unauthorized"
	if isUA(r, "curl") {
		unauthorized_message = `
It looks like you are trying to use the CURL api.
Here is a little documentation on how to do so!
------------------------------------------------------------
Shorten a URL:
  curl -su <username> --data "<url>" ` + host + `

Upload Text Snippet:
  curl -su <username> --data "<Text>" ` + host + `

  cat file.txt | curl -su <username> -d@- ` + host + `

Upload File:
  curl -su <username> --data-binary "@filename" ` + host + `

  cat file.txt | curl -su <username> --data-binary @- ` + host + `

`
	} else if isUA(r, "wget") {
		unauthorized_message = `
wget usage
------
wget -qO- --post-data "<url>"
`
	}
	w.WriteHeader(401)
	w.Write([]byte(unauthorized_message))
}

func webfinger(w http.ResponseWriter, r *http.Request) {
	//acct := r.URL.Query().Get("resource")
	type wfLink struct {
		Rel      string `json:"rel"`
		Type     string `json:"type"`
		Href     string `json:"href"`
		Template string `json:"href"`
	}
	type webfingerResponse struct {
		Subject string   `json:"subject"`
		Aliases []string `json:"aliases"`
		Links   []wfLink `json:"aliases"`
	}
	wf := webfingerResponse{
		Subject: "acct:demophoon@mastodon.brittg.com",
		Aliases: []string{
			"https://mastodon.brittg.com/@demophoon",
			"https://mastodon.brittg.com/users/demophoon",
		},
		Links: []wfLink{
			{
				Rel:  "http://webfinger.net/rel/profile-page",
				Type: "text/html",
				Href: "https://mastodon.brittg.com/@demophoon",
			},
			{
				Rel:  "self",
				Type: "application/activity+json",
				Href: "https://mastodon.brittg.com/users/demophoon",
			},
			{
				Rel:      "http://ostatus.org/schema/1.0/subscribe",
				Template: "https://mastodon.brittg.com/authorize_interaction?uri={uri}",
			},
		},
	}
	output, err := json.Marshal(wf)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(genericError))
		return
	}
	w.Write(output)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(healthcheckSuccess))
}

const (
	genericError       string = "An error occurred during the request"
	healthcheckSuccess string = "healthy"
)
