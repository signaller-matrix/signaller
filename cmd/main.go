package main

import (
	"log"

	internal "github.com/nxshock/signaller/internal"
)

var (
	server *internal.Server
)

func init() {
	server = internal.New()
	server.Address = "localhost"
	server.Backend = internal.NewMemoryBackend()
	server.Backend.Register("andrew", "1", "")
}

func main() {
	log.Println(server.Run())
}
