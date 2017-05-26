package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router    *mux.Router
	fileStore FileStore
	authToken string
}

func NewServer(fs FileStore, at string) *Server {
	s := &Server{
		router:    mux.NewRouter(),
		authToken: at,
	}

	s.buildRoutes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
