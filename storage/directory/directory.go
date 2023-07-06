package directorystate

import "gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"

type DirectoryStorage struct {
	upload_location string
}

func (d *DirectoryStorage) CreateFile(file []byte) (string, error) {
	return "", nil
}

func (d *DirectoryStorage) ReadFile(filepath string) ([]byte, error) {
	return []byte{}, nil
}

func (d *DirectoryStorage) DeleteFile(filepath string) error {
	return nil
}

func New(c *config.Config) *DirectoryStorage {
	return &DirectoryStorage{}
}
