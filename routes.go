package main

import "net/http"

func (s *Server) buildRoutes() {
	const (
		GET = "GET"
	)
	// Get to serve upload page
	s.router.Handle("/", http.HandlerFunc(s.getUploadPageHandler)).Methods(GET)
}
