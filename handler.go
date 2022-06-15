package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler to manage the home endpoint(/)
func HandlerRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Root endpoint working")
}

// Handler to manage the api endpoint
func HandlerHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API endpoint working")
}

// Handler to manage the post request in user endpoint
func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "aplication/json")
	w.Write(response)
}
