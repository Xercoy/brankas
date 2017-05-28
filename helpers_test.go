package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"strings"

	"github.com/cheekybits/is"
)

func TestFileUploadLimit(t *testing.T) {
	is := is.New(t)
	testImageFile := "./testdata/test-image.jpg"

	t.Run("detect image mimetype", func(t *testing.T) {
		testImage, err := os.Open(testImageFile)
		defer testImage.Close()
		is.NoErr(err)

		MimeTypeRdr := io.LimitReader(testImage, 512)
		mimeTypeBytes, err := ioutil.ReadAll(MimeTypeRdr)
		is.NoErr(err)

		mimetype := http.DetectContentType(mimeTypeBytes)
		is.True(strings.Contains(mimetype, "image"))
	})
}
