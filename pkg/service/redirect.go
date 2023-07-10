package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"
)

func (s *ShrlsService) Redirect(w http.ResponseWriter, r *http.Request) {
	shrl := r.URL.Path
	shrl = strings.TrimPrefix(shrl, "/")

	parts := strings.Split(shrl, ".")
	ref := &pb.Ref_ShortURL{
		Ref: &pb.Ref_ShortURL_Alias{
			Alias: parts[0],
		},
	}

	ctx := context.Background()
	redirect, err := s.NewClient().GetShrl(ctx, ref)
	if err != nil {
		if s.config.DefaultRedirect != "" {
			http.Redirect(w, r, s.config.DefaultRedirect, http.StatusTemporaryRedirect)
			log.Error("Unable to resolve Shrl. ", ref)
		} else {
			http.Error(w, fmt.Sprintf("Unable to fetch Shrl: %q. %s", ref, err), http.StatusNotFound)
		}
		return
	}

	switch redirect.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		http.Redirect(w, r, redirect.Content.GetUrl().Url, http.StatusTemporaryRedirect)

	case *pb.ExpandedURL_File:
		file := redirect.Content.GetFile()
		w.Write(file)

	case *pb.ExpandedURL_Snippet:
		w.Write([]byte(redirect.Content.GetSnippet().Body))
	}

	log.Debug(redirect)
}
