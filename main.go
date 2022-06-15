package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", "GET", HandlerRoot)
	server.Handle("/create", "POST", PostRequest)
	server.Handle("/user", "POST", UserPostRequest)
	server.Handle("/api", "POST", server.AddMiddleware(HandlerHome, CheckAuth(), Loggin()))
	server.Listen()
}
