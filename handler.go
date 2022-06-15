package main

import (
	"fmt"
	"net/http"
)

// Handle to manage the principal root ("/")
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root endpoint working")
}

// Handle to manage home
func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API endpoint working")
}
