package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func FileFromString(uploadPath string) URL {
	shrl := NewURL()
	shrl.UploadLocation = uploadPath
	shrl.Type = UploadedFile
	shrl.Create()
	return shrl
}

func URLFromFile(file []byte) (URL, error) {
	mtype := mimetype.Detect(file)

	err := os.MkdirAll(Settings.UploadDirectory, os.ModePerm)
	if err != nil {
		return URL{}, err
	}

	file_uuid_hyphen, err := uuid.New()
	if err != nil {
		return URL{}, err
	}
	str := hex.EncodeToString(file_uuid_hyphen[:])
	file_uuid := strings.Replace(str, "-", "", -1)
	filename := fmt.Sprintf("%s%s", file_uuid, mtype.Extension())
	filepath := path.Join(Settings.UploadDirectory, filename)

	err = ioutil.WriteFile(filepath, file, os.ModePerm)
	if err != nil {
		return URL{}, err
	}

	url := FileFromString(filepath)

	return url, nil
}

func fileUpload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url, err := URLFromFile(buf.Bytes())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	SuccessResponse(w, &url)
}
