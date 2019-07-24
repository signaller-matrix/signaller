package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/nxshock/signaller/internal/models/common"

	"github.com/nxshock/signaller/internal/models"
	login "github.com/nxshock/signaller/internal/models/login"
	register "github.com/nxshock/signaller/internal/models/register"
	mSync "github.com/nxshock/signaller/internal/models/sync"
	"github.com/nxshock/signaller/internal/models/versions"
	"github.com/nxshock/signaller/internal/models/whoami"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)
}

// https://models.org/docs/spec/client_server/latest#get-models-client-versions
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errorResponse(w, models.M_UNKNOWN, http.StatusBadRequest, "wrong method: "+r.Method)
		return
	}

	response := versions.Reply{Versions: []string{Version}}
	sendJsonResponse(w, http.StatusOK, response)
}

// https://models.org/docs/spec/client_server/latest#login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// https://models.org/docs/spec/client_server/latest#get-models-client-r0-login
	case "GET":
		{
			response := login.GetReply{
				Flows: []login.Flow{
					login.Flow{Type: common.AuthenticationTypePassword},
				},
			}

			sendJsonResponse(w, http.StatusOK, response)
		}
	// https://models.org/docs/spec/client_server/latest#post-models-client-r0-login
	case "POST":
		{
			var request login.PostRequest
			getRequest(r, &request) // TODO: handle error

			// delete start "@" if presents
			if strings.HasPrefix(request.Identifier.User, "@") {
				request.Identifier.User = strings.TrimPrefix(request.Identifier.User, "@")
			}

			token, apiErr := currServer.Backend.Login(request.Identifier.User, request.Password, request.DeviceID)
			if apiErr != nil {
				errorResponse(w, *apiErr, http.StatusForbidden, "")
				return
			}

			response := login.PostReply{
				UserID:      request.Identifier.User,
				AccessToken: token,
			}

			sendJsonResponse(w, http.StatusOK, response)
		}
	}
}

// https://models.org/docs/spec/client_server/latest#post-models-client-r0-logout
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, models.M_UNKNOWN, http.StatusBadRequest, "wrong method: "+r.Method)
		return
	}

	token := getTokenFromResponse(r)

	if token == "" {
		errorResponse(w, models.M_MISSING_TOKEN, http.StatusBadRequest, "")
		return
	}

	user := currServer.Backend.GetUserByToken(token)
	if user == nil {
		errorResponse(w, models.M_UNKNOWN_TOKEN, http.StatusBadRequest, "")
		return
	}

	user.Logout(token)

	sendJsonResponse(w, http.StatusOK, struct{}{})
}

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-logout-all
func LogoutAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, models.M_UNKNOWN, http.StatusBadRequest, "wrong method: "+r.Method)
		return
	}

	token := getTokenFromResponse(r)

	if token == "" {
		errorResponse(w, models.M_MISSING_TOKEN, http.StatusBadRequest, "")
		return
	}

	user := currServer.Backend.GetUserByToken(token)
	if user == nil {
		errorResponse(w, models.M_UNKNOWN_TOKEN, http.StatusBadRequest, "")
		return
	}

	user.LogoutAll()

	sendJsonResponse(w, http.StatusOK, struct{}{})
}

// https://models.org/docs/spec/client_server/latest#post-models-client-r0-register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errorResponse(w, models.M_UNKNOWN, http.StatusBadRequest, "wrong method: "+r.Method)
		return
	}

	kind := r.FormValue("kind")
	if kind != "user" {
		errorResponse(w, models.M_UNKNOWN, http.StatusBadRequest, "wrong kind: "+kind)
		return
	}

	var request register.RegisterRequest
	getRequest(r, &request) // TODO: handle error

	_, token, apiErr := currServer.Backend.Register(request.Username, request.Password, request.DeviceID)
	if apiErr != nil {
		errorResponse(w, *apiErr, http.StatusBadRequest, "")
		return
	}

	var response register.RegisterResponse
	response.UserID = "@" + request.Username
	response.DeviceID = request.DeviceID
	response.AccessToken = token

	sendJsonResponse(w, http.StatusOK, response)
}

// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-account-whoami
func WhoAmIHandler(w http.ResponseWriter, r *http.Request) {
	token := getTokenFromResponse(r)
	if token == "" {
		errorResponse(w, models.M_FORBIDDEN, http.StatusForbidden, "")
	}

	user := currServer.Backend.GetUserByToken(token)

	response := whoami.Response{UserID: user.ID()}

	sendJsonResponse(w, http.StatusOK, response)
}

// https://models.org/docs/spec/client_server/latest#get-models-client-r0-sync
func SyncHandler(w http.ResponseWriter, r *http.Request) {
	var request mSync.SyncRequest
	request.Filter = r.FormValue("filter")

	timeout, err := strconv.Atoi(r.FormValue("timeout"))
	if err != nil {
		errorResponse(w, models.M_UNKNOWN, http.StatusBadRequest, "timeout parse failes")
		return
	}
	request.Timeout = timeout
	//log.Printf("%+v", request)

	token := getTokenFromResponse(r)
	if token == "" {
		errorResponse(w, models.M_MISSING_TOKEN, http.StatusBadRequest, "")
		return
	}

	response, _ := currServer.Backend.Sync(token, request) // TODO: handle error

	response.NextBatch = "123"
	response.Rooms = mSync.RoomsSyncReply{}

	sendJsonResponse(w, http.StatusOK, response)
}

func sendJsonResponse(w http.ResponseWriter, httpStatus int, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, err = w.Write(b) // TODO: handle n
	return err
}

func getRequest(r *http.Request, request interface{}) error {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, request)
}
