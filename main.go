package main

func main() {
	server := NewServer(":3000")
	server.Handle("/", HandlerRoot)
	server.Handle("/home", server.AddMiddleware(HandlerHome, CheckAuth(), Loggin()))
	server.Listen()
}
