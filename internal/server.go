package internal

import (
	"net/http"

	"github.com/gorilla/mux"

	"errors"

	"strconv"

	"github.com/nxshock/signaller/internal/models/capabilities"
)

var currServer *Server

type Server struct {
	httpServer *http.Server
	router     *mux.Router

	Address string

	Capabilities capabilities.Capabilities
	Backend      Backend
}

func NewServer(port int) (*Server, error) {

	router := mux.NewRouter()
	router.HandleFunc("/_matrix/client/versions", VersionHandler)
	router.HandleFunc("/_matrix/client/r0/login", LoginHandler)
	router.HandleFunc("/_matrix/client/r0/logout", LogoutHandler)
	router.HandleFunc("/_matrix/client/r0/logout/all", LogoutAllHandler)
	router.HandleFunc("/_matrix/client/r0/register", RegisterHandler)
	router.HandleFunc("/_matrix/client/r0/account/whoami", WhoAmIHandler)
	router.HandleFunc("/_matrix/client/r0/joined_rooms", JoinedRoomsHandler)
	router.HandleFunc("/_matrix/client/r0/account/password", PasswordHandler)
	router.HandleFunc("/_matrix/client/r0/sync", SyncHandler)
	router.HandleFunc("/_matrix/client/r0/capabilities", CapabilitiesHandler)
	router.HandleFunc("/_matrix/client/r0/devices", DevicesHandler)
	router.HandleFunc("/", RootHandler)

	if port <= 0 || port > 65535 {
		return nil, errors.New("invalid port number")
	}

	httpServer := new(http.Server)
	httpServer.Addr = ":" + strconv.Itoa(port)
	httpServer.Handler = router

	server := &Server{
		httpServer: httpServer,
		router:     router}

	currServer = server
	return server, nil
}

func (server *Server) Run() error {
	return server.httpServer.ListenAndServe() // TODO: custom port
}
