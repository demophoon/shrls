package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func FileFromString(uploadPath string) URL {
	shrl := NewURL()
	shrl.UploadLocation = uploadPath
	shrl.Type = UploadedFile
	createUrl(&shrl)
	return shrl
}

func fileUpload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	err = os.MkdirAll(Settings.UploadDirectory, os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file_uuid_hyphen, err := uuid.New()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	str := hex.EncodeToString(file_uuid_hyphen[:])
	file_uuid := strings.Replace(str, "-", "", -1)
	filename := fmt.Sprintf("%s%s", file_uuid, filepath.Ext(fileHeader.Filename))
	filepath := path.Join(Settings.UploadDirectory, filename)
	dst, err := os.Create(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	url := FileFromString(filepath)

	SuccessResponse(w, &url)
}
