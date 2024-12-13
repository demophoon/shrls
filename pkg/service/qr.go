package service

import (
	"net/url"

	pb "github.com/demophoon/shrls/server/gen"

	qrcode "github.com/skip2/go-qrcode"
)

func (s *ShrlsService) toEncodedQR(shrl *pb.ShortURL) (*qrcode.QRCode, error) {
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

	qr, err := qrcode.New(location, qrcode.Medium)
	if err != nil {
		return nil, err
	}

	return qr, nil
}

func (s *ShrlsService) ToQR(shrl *pb.ShortURL) ([]byte, error) {
	qr, err := s.toEncodedQR(shrl)
	if err != nil {
		return nil, err
	}

	return qr.PNG(256)
}

func (s *ShrlsService) ToTextQR(shrl *pb.ShortURL) ([]byte, error) {
	qr, err := s.toEncodedQR(shrl)
	if err != nil {
		return nil, err
	}

	return []byte(qr.ToSmallString(false)), nil
}
