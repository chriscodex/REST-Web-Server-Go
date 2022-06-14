package main

import (
	"fmt"
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

// Methods

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

// This function turns Router into a Handler(Interface)
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Server Working")
}
