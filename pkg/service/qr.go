package service

import (
	pb "gitlab.cascadia.demophoon.com/demophoon/go-shrls/server/gen"

	qrcode "github.com/skip2/go-qrcode"
)

func (s *ShrlsService) ToQR(url *pb.ShortURL) ([]byte, error) {
	location := s.config.BaseURL + "/" + url.Stub

	var result []byte
	result, err := qrcode.Encode(location, qrcode.Medium, 256)

	if err != nil {
		return nil, err
	}

	return result, nil
}
