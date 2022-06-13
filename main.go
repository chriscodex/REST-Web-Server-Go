package main

import "fmt"

func main() {
	server := NewServer(":3000")
	//server.Listen()
	fmt.Println(server)
}
