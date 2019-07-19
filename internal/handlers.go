package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/nxshock/signaller/internal/models"
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

	response := models.VersionsReply{Versions: []string{Version}}
	sendJsonResponse(w, http.StatusOK, response)
}

// https://models.org/docs/spec/client_server/latest#login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// https://models.org/docs/spec/client_server/latest#get-models-client-r0-login
	case "GET":
		{
			type LoginFlow struct {
				Type string `json:"type"`
			}

			type Response struct {
				Flows []LoginFlow `json:"flows"`
			}

			response := Response{
				Flows: []LoginFlow{
					LoginFlow{Type: string(M_LOGIN_PASSWORD)}}}

			sendJsonResponse(w, http.StatusOK, response)
		}

	// https://models.org/docs/spec/client_server/latest#post-models-client-r0-login
	case "POST":
		{
			var request models.LoginRequest
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

			response := models.LoginReply{
				UserID:      request.Identifier.User,
				AccessToken: token}

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

	apiErr := currServer.Backend.Logout(token)
	if apiErr != nil {
		errorResponse(w, *apiErr, http.StatusBadRequest, "") // TODO: check code
		return
	}

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

	var request models.RegisterRequest
	getRequest(r, &request) // TODO: handle error

	token, apiErr := currServer.Backend.Register(request.Username, request.Password, request.DeviceID)
	if apiErr != nil {
		errorResponse(w, *apiErr, http.StatusBadRequest, "")
		return
	}

	var response models.RegisterResponse
	response.UserID = "@" + request.Username
	response.DeviceID = request.DeviceID
	response.AccessToken = token

	sendJsonResponse(w, http.StatusOK, response)
}

// https://models.org/docs/spec/client_server/latest#get-models-client-r0-sync
func SyncHandler(w http.ResponseWriter, r *http.Request) {
	var request models.SyncRequest
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
	response.Rooms = models.RoomsSyncReply{}

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
