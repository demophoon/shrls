package service

import (
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/mat/besticon/v3/besticon"
	log "github.com/sirupsen/logrus"

	pb "github.com/demophoon/shrls/server/gen"
)

func (s Server) enhanceShortUrl(u *pb.ShortURL) {
	// Fetch Favicon
	s.resolveFavicon(u)

	// Resolve Redirects
	s.resolveRedirects(u)

	// Strip Url Parameters
	s.removeQueryParameters(u)
}

func (s Server) resolveRedirects(u *pb.ShortURL) {
	switch u.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		if s.config.ResolveURLMatchingHosts == nil {
			return
		}

		loc := u.Content.GetUrl().Url
		parsed, err := url.Parse(loc)
		if err != nil {
			log.Warnf("Couldn't parse url when resolving redirects from %s: %s", loc, err)
			return
		}

		s := *s.config.ResolveURLMatchingHosts
		regex_str := strings.Join(s, "|")
		log.Debugf("Regexp string: %s", regex_str)
		matcher, err := regexp.Compile(regex_str)
		if err != nil {
			log.Errorf("Unable to determine resolve_urls_matching_hosts option: %s", err)
			return
		}
		if !matcher.Match([]byte(parsed.Host)) {
			return
		}

		nextLocation := parsed.String()
		i := 0
		for i < 25 {
			client := &http.Client{
				CheckRedirect: func(req *http.Request, via []*http.Request) error {
					return http.ErrUseLastResponse
				},
			}

			resp, err := client.Head(nextLocation)
			if err != nil {
				log.Errorf("Unable to resolve url: %s", err)
				return
			}

			if resp.StatusCode >= 300 && resp.StatusCode < 400 {
				nextLocation = resp.Header.Get("Location")
				i += 1
			} else {
				break
			}
		}

		u.Content.Content = &pb.ExpandedURL_Url{
			Url: &pb.Redirect{
				Url:     nextLocation,
				Favicon: u.Content.GetUrl().Favicon,
			},
		}

	}
}

func (s Server) removeQueryParameters(u *pb.ShortURL) {
	switch u.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		if s.config.RemoveQueryParametersMatchingHosts == nil {
			return
		}

		loc := u.Content.GetUrl().Url
		parsed, err := url.Parse(loc)
		if err != nil {
			log.Warnf("Couldn't parse url when removing query parameters from %s: %s", loc, err)
			return
		}

		s := *s.config.RemoveQueryParametersMatchingHosts
		regex_str := strings.Join(s, "|")
		log.Debugf("Regexp string: %s", regex_str)
		matcher, err := regexp.Compile(regex_str)
		if err != nil {
			log.Errorf("Unable to determine remove_query_parameters_matching_hosts option: %s", err)
			return
		}
		if !matcher.Match([]byte(parsed.Host)) {
			return
		}

		stripped := url.URL{
			Scheme: parsed.Scheme,
			User:   parsed.User,
			Host:   parsed.Host,
			Path:   parsed.Path,
		}
		u.Content.Content = &pb.ExpandedURL_Url{
			Url: &pb.Redirect{
				Url:     stripped.String(),
				Favicon: u.Content.GetUrl().Favicon,
			},
		}
	}
}

func (s Server) resolveFavicon(u *pb.ShortURL) {
	switch u.Content.Content.(type) {
	case *pb.ExpandedURL_Url:
		loc := u.Content.GetUrl().Url

		bi := besticon.New()
		icf := bi.NewIconFinder()

		icons, err := icf.FetchIcons(loc)

		if err == nil {
			for _, icon := range icons {
				if icon.Width >= 16 && icon.Height >= 16 {
					u.Content.Content = &pb.ExpandedURL_Url{
						Url: &pb.Redirect{
							Url:     loc,
							Favicon: icon.ImageData,
						},
					}
					log.Tracef("favicon found for %s at %s", loc, icon.URL)
					break
				}
			}
		} else {
			log.Warnf("Unable to fetch favicon from %s: %s", loc, err)
		}
	}
}
