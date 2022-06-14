package main

import (
	"fmt"
	"net/http"
)

// Handle to manage the principal root ("/")
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root working")
}
