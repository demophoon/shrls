package directorystate

import (
	"os"
	"path"

	"gitlab.cascadia.demophoon.com/demophoon/go-shrls/pkg/config"

	"github.com/google/uuid"
)

type DirectoryStorage struct {
	upload_location string
}

func (d *DirectoryStorage) CreateFile(file []byte) (string, error) {
	err := os.MkdirAll(d.upload_location, os.ModePerm)
	if err != nil {
		return "", err
	}

	key := uuid.New()
	if err != nil {
		return "", err
	}
	filepath := path.Join(d.upload_location, key.String())

	err = os.WriteFile(filepath, file, os.ModePerm)
	if err != nil {
		return "", err
	}

	return key.String(), nil
}

func (d *DirectoryStorage) ReadFile(key string) ([]byte, error) {
	filepath := path.Join(d.upload_location, key)
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (d *DirectoryStorage) DeleteFile(key string) error {
	filepath := path.Join(d.upload_location, key)
	return os.Remove(filepath)
}

func New(c *config.Config) *DirectoryStorage {
	storage := &DirectoryStorage{}
	storage.upload_location = c.UploadDirectory
	return storage
}
