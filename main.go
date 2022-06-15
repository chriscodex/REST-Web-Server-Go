package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", "GET", HandlerRoot)
	server.Handle("/home", "POST", server.AddMiddleware(HandlerHome, CheckAuth(), Loggin()))
	server.Listen()
}
