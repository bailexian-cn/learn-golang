package upload

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

var cacheDir = "/tmp"

func Server() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			uploadHandle(w, r)
		default:
			SendResponse(w, http.StatusNotFound, Response{"Method not support."})
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func SendResponse(w http.ResponseWriter, code int, resp interface{}) {
	enc, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Failed JSON-encoding HTTP response: %v", err)
		return
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(enc)
	if err != nil {
		log.Printf("Failed sending HTTP response body: %v", err)
	}
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	mFile, fileHeader, err := r.FormFile("chart")
	defer func(file multipart.File) {
		_ = file.Close()
	}(mFile)
	if err != nil {
		SendResponse(w, http.StatusBadRequest, Response{"Can't find param chart."})
		return
	}
	fileName := fileHeader.Filename
	err = cacheFile(fileName, mFile)
	if err != nil {
		SendResponse(w, http.StatusInternalServerError, Response{"Cache file err."})
		return
	}
	log.Printf("Success to cache file %s.", fileName)
	w.WriteHeader(http.StatusCreated)
}

func cacheFile(name string, r io.Reader) error {
	absolutePath := fmt.Sprintf("%s/%s", cacheDir, name)
	file, err := os.Create(absolutePath)
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	if err != nil {
		return err
	}
	var buf [1024]byte
	for {
		n, err := r.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		m, err := file.Write(buf[:n])
		if m != n {
			return errors.New("cache file error")
		}
		if err != nil {
			return err
		}
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}
