package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

func uploadPageTemplate(authToken, fileName string, w io.Writer) error {
	tmplFile, err := os.Open("form.html")
	defer tmplFile.Close()
	if err != nil {
		return err
	}

	tmplBytes, err := ioutil.ReadAll(tmplFile)
	if err != nil {
		return err
	}

	auth := struct {
		AuthToken string
	}{
		AuthToken: authToken,
	}

	t := template.New("uploadPage")
	t.Parse(string(tmplBytes))
	t.Execute(w, auth)

	return nil
}
