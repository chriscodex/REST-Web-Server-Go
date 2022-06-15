package main

import "net/http"

// Struct of middleware
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Struct of User
type User struct {
	Name  string
	Email string
	Phone string
}
