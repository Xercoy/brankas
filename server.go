package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router        *mux.Router
	fileStore     FileStore
	authToken     string
	fileByteLimit int64
}

func NewServer(fs FileStore, at string, ful int64) *Server {
	s := &Server{
		router:        mux.NewRouter(),
		authToken:     at,
		fileByteLimit: ful,
	}

	s.buildRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
