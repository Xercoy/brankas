package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/cheekybits/is"
)

// Test variables: tokens, file type, file size
func TestPostUploadPageHandler(t *testing.T) {
	var fileStore FileStore
	is := is.New(t)

	// Server setup.
	s := NewServer(fileStore, "PostUploadToken", 8388608)
	routeHandler := http.HandlerFunc(s.postUploadPageHandler)
	testServer := httptest.NewServer(routeHandler)
	defer testServer.Close()

	response, respBody := postUploadPage(t, testServer.URL, s.authToken)
	is.Equal(response.StatusCode, 200)

	log.Printf("Response Body:\n%s", respBody)
}

// Test variables: page and token
func TestGetUploadPageHandler(t *testing.T) {

	var fileStore FileStore
	is := is.New(t)

	// Server setup.
	s := NewServer(fileStore, "TESTTOKEN", 8388608)
	routeHandler := http.HandlerFunc(s.getUploadPageHandler)
	testServer := httptest.NewServer(routeHandler)
	defer testServer.Close()

	// Need page output here to assert.
	assertBody := new(bytes.Buffer)
	is.NoErr(uploadPageTemplate(s.authToken, "form.html", assertBody))

	t.Run("correctly receive upload page", func(t *testing.T) {
		resp, respBody := getUploadPage(t, testServer.URL)
		is.Equal(resp.StatusCode, 200)
		is.Equal(respBody, assertBody.String())
	})
}

func postUploadPage(t *testing.T, URL string, authToken string) (*http.Response, string) {
	is := is.New(t)
	file := "./testdata/test-image.jpg"

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	uploadFile, err := os.Open(file)
	defer uploadFile.Close()

	is.NoErr(err)

	formFile, err := w.CreateFormFile("data", file)
	is.NoErr(err)

	formFile, err = w.CreateFormField("auth")
	is.NoErr(err)

	bytesWritten, err := formFile.Write([]byte(authToken))
	is.NotEqual(bytesWritten, 0)
	is.NoErr(err)

	formBytesWritten, err := io.Copy(formFile, uploadFile)
	log.Printf("%d bytes written to uploadFile from form\n", formBytesWritten)
	is.NoErr(err)

	w.Close()
	log.Printf("Size of buffer is %d", len(b.Bytes()))
	request, err := http.NewRequest("POST", URL, &b)
	is.NoErr(err)

	request.Header.Set("Content-Type", w.FormDataContentType())

	/* Debug
	reqBody, err := ioutil.ReadAll(request.Body)
	log.Printf("Request Body: %s", reqBody) */

	client := &http.Client{}
	response, err := client.Do(request)
	is.NoErr(err)

	respBody, err := ioutil.ReadAll(response.Body)
	is.NoErr(err)

	return response, string(respBody)
}

func getUploadPage(t *testing.T, URL string) (*http.Response, string) {
	is := is.New(t)

	request, err := http.NewRequest("GET", URL, nil)
	is.NoErr(err)

	client := &http.Client{}
	response, err := client.Do(request)
	is.NoErr(err)

	bodyInBytes, err := ioutil.ReadAll(response.Body)
	is.NoErr(err)

	return response, string(bodyInBytes)
}
