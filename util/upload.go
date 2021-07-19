package util

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fileName := handler.Filename
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f, err := os.OpenFile("./public/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.Copy(f, file)
	json.NewEncoder(w).Encode(map[string]string{
		"image_url": "http://localhost:3000/images/" + fileName,
	})
}
