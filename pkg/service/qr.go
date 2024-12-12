package service

import (
	"net/url"

	pb "github.com/demophoon/shrls/server/gen"

	qrcode "github.com/skip2/go-qrcode"
)

func (s *ShrlsService) ToQR(shrl *pb.ShortURL) ([]byte, error) {
	url_scheme := "http"
	if s.config.DefaultRedirectSsl {
		url_scheme = "https"
	}
	u := url.URL{
		Scheme: url_scheme,
		Host:   s.config.Host,
		Path:   shrl.Stub,
	}

	location := u.String()

	var result []byte
	result, err := qrcode.Encode(location, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}

	return result, nil
}
