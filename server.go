package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	httpServer *http.Server
	router     *mux.Router

	Address string
	Backend Backend
}

func New() *Server {
	router := mux.NewRouter()
	router.HandleFunc("/_matrix/client/versions", VersionHandler)
	router.HandleFunc("/_matrix/client/r0/login", LoginHandler)
	router.HandleFunc("/_matrix/client/r0/logout", LogoutHandler)
	router.HandleFunc("/_matrix/client/r0/register", RegisterHandler)
	router.HandleFunc("/_matrix/client/r0/sync", SyncHandler)
	router.HandleFunc("/", RootHandler)

	httpServer := new(http.Server)
	httpServer.Addr = ":80"
	httpServer.Handler = router

	server := &Server{
		httpServer: httpServer,
		router:     router}

	return server
}

func (server *Server) Run() error {
	return server.httpServer.ListenAndServe() // TODO: custom port
}
