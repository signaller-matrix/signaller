package main

import (
	"log"
)

var (
	server *Server
)

func init() {
	server = New()
	server.Address = "localhost"
	server.Backend = NewMemoryBackend()
	server.Backend.Register("andrew", "1", "")
}

func main() {
	log.Println(server.Run())
}
