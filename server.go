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
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	s.router.rules[path][method] = handler
}

// Function to add a middleware
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
