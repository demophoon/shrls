package directorystate

import (
	"os"
	"path"

	"github.com/demophoon/shrls/pkg/config"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type DirectoryStorage struct {
	upload_location string
}

func (d *DirectoryStorage) CreateFile(file []byte) (string, int64, error) {
	err := os.MkdirAll(d.upload_location, os.ModePerm)
	if err != nil {
		return "", 0, err
	}

	key := uuid.New()
	filepath := path.Join(d.upload_location, key.String())

	err = os.WriteFile(filepath, file, os.ModePerm)
	if err != nil {
		return "", 0, err
	}

	return key.String(), int64(len(file)), nil
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
	if c.UploadBackend.Directory == nil {
		log.Fatal("Couldn't initialize upload backend. Path not defined.")
	}
	storage := &DirectoryStorage{}
	storage.upload_location = c.UploadBackend.Directory.Path
	return storage
}
