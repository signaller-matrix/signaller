package memory

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/nxshock/signaller/internal"
	"github.com/nxshock/signaller/internal/models"
	"github.com/nxshock/signaller/internal/models/common"
	mSync "github.com/nxshock/signaller/internal/models/sync"
)

var first bool

type Backend struct {
	data  map[string]*User
	mutex sync.Mutex // TODO: replace with RW mutex
}

type User struct {
	Password string
	Tokens   map[string]Token
}

type Token struct {
	Device string
}

func NewBackend() *Backend {
	return &Backend{data: make(map[string]*User)}
}

func (backend Backend) Register(username, password, device string) (token string, err *models.ApiError) {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	if _, ok := backend.data[username]; ok {
		return "", internal.NewError(models.M_USER_IN_USE, "trying to register a user ID which has been taken")
	}

	token = internal.NewToken(internal.DefaultTokenSize)

	backend.data[username] = &User{
		Password: password,
		Tokens: map[string]Token{
			token: {
				Device: device}}}

	return token, nil
}

func (backend Backend) Login(username, password, device string) (token string, err *models.ApiError) {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	user, ok := backend.data[username]
	if !ok {
		return "", internal.NewError(models.M_FORBIDDEN, "wrong username")
	}

	if user.Password != password {
		return "", internal.NewError(models.M_FORBIDDEN, "wrong password")
	}

	token = internal.NewToken(internal.DefaultTokenSize)

	backend.data[username].Tokens[token] = Token{Device: device}

	return token, nil
}

func (backend Backend) Logout(token string) *models.ApiError {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	for _, user := range backend.data {
		for userToken, _ := range user.Tokens {
			if userToken == token {
				delete(user.Tokens, token)
				return nil
			}
		}
	}

	return internal.NewError(models.M_UNKNOWN_TOKEN, "unknown token") // TODO: create error struct
}

func (backend Backend) Sync(token string, request mSync.SyncRequest) (response *mSync.SyncReply, err *models.ApiError) {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	log.Println(request)

	if !first {
		log.Println(1)
		response = &mSync.SyncReply{
			AccountData: common.AccountData{
				Events: []common.Event{
					common.Event{Type: "m.direct", Content: json.RawMessage(`"@vasyo2:localhost":"!room1:localhost"`)},
				}},
			Rooms: mSync.RoomsSyncReply{
				Join: map[string]common.JoinedRoom{
					"!room1:localhost": common.JoinedRoom{
						Timeline: common.Timeline{
							Events: []common.RoomEvent{
								common.RoomEvent{Type: "m.room.create", Sender: "@vasyo2:localhost"},
								common.RoomEvent{Type: "m.room.member", Sender: "@vasyo2:localhost", Content: json.RawMessage(`membership:"join",displayname:"vasyo2"`)},
							}}}}}}
		/*					InviteState: models.InviteState{
							Events: []models.StrippedState{
								models.StrippedState{Type: "m.room.join_rules", Content: json.RawMessage(`join_rule:"invite"`), Sender: "@vasyo2:" + server.Address},
								models.StrippedState{Type: "m.room.member", Content: json.RawMessage(`membership:"join",displayname:"vasyo2"`), Sender: "@vasyo2:" + server.Address},
								models.StrippedState{Type: "m.room.member", Content: json.RawMessage(`is_direct:"true",membership:"invite",displayname:"vasyo"`), Sender: "@vasyo2:" + server.Address},
							}}}}}}*/
		first = true
	} else {
		os.Exit(0)
		response = &mSync.SyncReply{}
	}

	return response, nil // TODO: implement
}
