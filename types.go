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
	Name  string
	Email string
	Phone string
}

// Method to parse his properties in json
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
