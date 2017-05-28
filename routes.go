package main

import "net/http"

func (s *Server) buildRoutes() {
	const (
		GET  = "GET"
		POST = "POST"
	)
	// Get to serve upload page
	s.router.Handle("/upload", http.HandlerFunc(s.getUploadPageHandler)).Methods(GET)

	// Post images to the upload route
	s.router.Handle("/upload", http.HandlerFunc(s.postUploadPageHandler)).Methods(POST)
}
