package main

import "net/http"

// Struct of Server
type Server struct {
	port   string
	router *Router
}

// Function to create a new server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Methods
// Function to listen with listen and serve
func (s *Server) Listen() error {
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}

// Function to add a path to a handler
func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}
