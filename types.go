package main

import (
	"encoding/json"
	"net/http"
)

// Struct of middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Struct of Metadata
type Metadata interface {
}

// Struct of User
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Method to parse his properties in json
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
