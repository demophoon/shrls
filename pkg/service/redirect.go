package service

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	pb "github.com/demophoon/shrls/server/gen"
	log "github.com/sirupsen/logrus"
)

func isUA(r *http.Request, agent string) bool {
	ua := strings.ToLower(r.UserAgent())
	return strings.HasPrefix(ua, strings.ToLower(agent))
}

func isTerminal(r *http.Request) bool {
	return isUA(r, "curl") || isUA(r, "wget")
}

func (s *ShrlsService) Redirect(w http.ResponseWriter, r *http.Request) {
	shrl := r.URL.Path
	shrl = strings.TrimPrefix(shrl, "/")

	parts := strings.Split(shrl, ".")

	var ext string
	if len(parts) > 1 {
		ext = parts[1]
	}

	ref := &pb.Ref_ShortURL{
		Ref: &pb.Ref_ShortURL_Alias{
			Alias: parts[0],
		},
	}

	ctx := context.Background()
	redirect, err := s.NewClient().GetShrl(ctx, ref)
	if err != nil {
		if isTerminal(r) && s.config.DefaultTerminalString != "" {
			w.Write([]byte(s.config.DefaultTerminalString))
			return
		}
		if s.config.DefaultRedirect != "" {
			http.Redirect(w, r, s.config.DefaultRedirect, http.StatusTemporaryRedirect)
			log.Error("Unable to resolve Shrl. ", ref)
		} else {
			http.Error(w, fmt.Sprintf("Unable to fetch Shrl: %q. %s", ref, err), http.StatusNotFound)
		}
		return
	}

	switch ext {
	case "qr", "qrcode":
		var qr []byte
		var err error
		if isTerminal(r) {
			qr, err = s.ToTextQR(redirect)
		} else {
			qr, err = s.ToQR(redirect)
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to convert to QR: %s", err), http.StatusInternalServerError)
			return
		}
		w.Write(qr)
	default:
		switch redirect.Content.Content.(type) {
		case *pb.ExpandedURL_Url:
			http.Redirect(w, r, redirect.Content.GetUrl().Url, http.StatusTemporaryRedirect)

		case *pb.ExpandedURL_File:
			key := redirect.Content.GetFile().Ref
			f, err := s.storage.ReadFile(key)
			if err != nil {
				http.Error(w, "Unable to locate file", http.StatusInternalServerError)
			}
			w.Write(f)

		case *pb.ExpandedURL_Snippet:
			w.Write([]byte(redirect.Content.GetSnippet().Body))
		}
	}

	log.Debug(redirect)
}
