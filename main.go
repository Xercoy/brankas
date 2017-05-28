package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	var fs FileStore
	var authTokenFlag, addressFlag string
	var fileUploadLimitFlag int64

	flag.Int64Var(&fileUploadLimitFlag, "file-limit", 8388608, "int64 value of size in bytes of file limit.")
	flag.StringVar(&addressFlag, "address", ":8049", "address for server, default is localhost, port 8049.")
	flag.StringVar(&authTokenFlag, "auth-token", "", "auth token, can also be set with environment variable AUTHTOKEN.")
	flag.Parse()

	if authTokenFlag == "" {
		authTokenFlag = os.Getenv("AUTHTOKEN")
	}

	if authTokenFlag == "" {
		log.Fatal(errors.New("error: AUTHTOKEN env variable or auth-token flag must be set"))
	}

	s := NewServer(fs, authTokenFlag, fileUploadLimitFlag)
	server := &http.Server{
		Addr:    addressFlag,
		Handler: s,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
