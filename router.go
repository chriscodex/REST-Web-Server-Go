package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// Methods

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Function to map a handler with a path and a method http
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, exist, methodExist
}

// Function to assign the response of the server
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exist, methodExist := r.FindHandler(request.URL.Path, request.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
