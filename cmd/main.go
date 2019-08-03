package main

import (
	"log"
	"strconv"

	"github.com/signaller-matrix/signaller/internal"
	"github.com/signaller-matrix/signaller/internal/backends/memory"
)

var (
	server            *internal.Server
	defaultPortNumber = 8008
)

func init() {
	var err error
	server, err = internal.NewServer(defaultPortNumber)
	if err != nil {
		panic(err)
	}
	server.Address = "localhost"
	server.Backend = memory.NewBackend(server.Address)
	server.Backend.Register("andrew", "1", "")
}

func main() {
	log.Println("Server started on port " + strconv.Itoa(defaultPortNumber))
	log.Println(server.Run())
}
