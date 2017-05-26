package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cheekybits/is"
)

// Test variables: page and token
func TestGetUploadPageHandler(t *testing.T) {

	var fileStore FileStore
	is := is.New(t)

	// Server setup.
	s := NewServer(fileStore, "TESTTOKEN")
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

// Test variables: tokens, file type, file size
func TestPostUploadPageHandler(t *testing.T) {

}
