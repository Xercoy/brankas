package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (s *Server) getUploadPageHandler(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	err := uploadPageTemplate(s.authToken, "form.html", buf)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", buf.String())
}
