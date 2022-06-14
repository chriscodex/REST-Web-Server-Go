package main

import "net/http"

type Router struct {
	rules map[string]http.HandlerFunc
}
