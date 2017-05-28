package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Check auth token, file size, check metadata
func (s *Server) postUploadPageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10000000) //refactor
	if err != nil {
		fmt.Printf("Multipart error: %v", err)
		return
	}

	// Token is either not provided or there's a mismatch.
	authToken := r.PostFormValue("auth")
	if authToken == "" || authToken != s.authToken {
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Printf("Post Upload Handler: auth is empty.\nBody:%s\n", body) // TODO http.Error
		http.Error(w, "Authorization token incorrect or missing.", http.StatusForbidden)
		return
	}

	// Read the first 512 bytes from the body.
	fileHeaderRdr := http.MaxBytesReader(w, r.Body, 512)
	defer fileHeaderRdr.Close()
	fileHeaderBytes, err := ioutil.ReadAll(fileHeaderRdr)
	if err != nil {
		panic(err) // TODO
	}

	// Retrieve mime type and error out if it's not an image based file.
	mimeType := detectMimeType(fileHeaderBytes)
	if strings.Contains(mimeType, "image") == false {
		http.Error(w, "Image upload mime type must be of type image", http.StatusForbidden)
		panic(err) // TODO
	}

	// Limit the amount being read from the body.
	fileRdr := http.MaxBytesReader(w, r.Body, 10000000)
	fileBytes, err := ioutil.ReadAll(fileRdr)
	if err != nil {
		panic(err)
	}

	// "stitch" files back together
	completeFile := append(fileHeaderBytes, fileBytes...)

	// file uploaded is too large.
	if int64(len(completeFile)) > s.fileByteLimit {
		http.Error(w, "Upload image file size too large", http.StatusForbidden)
	}

	// Download file to temp folder/
	// Place file into database.
}

// Retrieve first 512 bytes and detect mimetype.
func detectMimeType(data []byte) string {
	return http.DetectContentType(data)
}

func (s *Server) getUploadPageHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	err := uploadPageTemplate(s.authToken, "form.html", buf)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", buf.String())
}
