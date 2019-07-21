package main

import (
	"log"

	"github.com/nxshock/signaller/internal"
	"github.com/nxshock/signaller/internal/backends/memory"
)

var (
	server *internal.Server
)

func init() {
	server = internal.New()
	server.Address = "localhost"
	server.Backend = memory.NewBackend(server.Address)
	server.Backend.Register("andrew", "1", "")
}

func main() {
	log.Println(server.Run())
}
