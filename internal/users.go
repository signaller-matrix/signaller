package internal

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/nxshock/signaller/internal/models"
)

var first bool

type MemoryBackend struct {
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

func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{data: make(map[string]*User)}
}

func (memoryBackend MemoryBackend) Register(username, password, device string) (token string, err *models.ApiError) {
	memoryBackend.mutex.Lock()
	defer memoryBackend.mutex.Unlock()

	if _, ok := memoryBackend.data[username]; ok {
		return "", NewError(models.M_USER_IN_USE, "trying to register a user ID which has been taken")
	}

	token = newToken(DefaultTokenSize)

	memoryBackend.data[username] = &User{
		Password: password,
		Tokens: map[string]Token{
			token: {
				Device: device}}}

	return token, nil
}

func (memoryBackend MemoryBackend) Login(username, password, device string) (token string, err *models.ApiError) {
	memoryBackend.mutex.Lock()
	defer memoryBackend.mutex.Unlock()

	user, ok := memoryBackend.data[username]
	if !ok {
		return "", NewError(models.M_FORBIDDEN, "wrong username")
	}

	if user.Password != password {
		return "", NewError(models.M_FORBIDDEN, "wrong password")
	}

	token = newToken(DefaultTokenSize)

	memoryBackend.data[username].Tokens[token] = Token{Device: device}

	return token, nil
}

func (memoryBackend MemoryBackend) Logout(token string) *models.ApiError {
	memoryBackend.mutex.Lock()
	defer memoryBackend.mutex.Unlock()

	for _, user := range memoryBackend.data {
		for userToken, _ := range user.Tokens {
			if userToken == token {
				delete(user.Tokens, token)
				return nil
			}
		}
	}

	return NewError(models.M_UNKNOWN_TOKEN, "unknown token") // TODO: create error struct
}

func (memoryBackend MemoryBackend) Sync(token string, request models.SyncRequest) (response *models.SyncReply, err *models.ApiError) {
	memoryBackend.mutex.Lock()
	defer memoryBackend.mutex.Unlock()

	log.Println(request)

	if !first {
		log.Println(1)
		response = &models.SyncReply{
			AccountData: models.AccountData{
				Events: []models.Event{
					models.Event{Type: "m.direct", Content: json.RawMessage(`"@vasyo2:localhost":"!room1:localhost"`)},
				}},
			Rooms: models.RoomsSyncReply{
				Join: map[string]models.JoinedRoom{
					"!room1:localhost": models.JoinedRoom{
						Timeline: models.Timeline{
							Events: []models.RoomEvent{
								models.RoomEvent{Type: "m.room.create", Sender: "@vasyo2:localhost"},
								models.RoomEvent{Type: "m.room.member", Sender: "@vasyo2:localhost", Content: json.RawMessage(`membership:"join",displayname:"vasyo2"`)},
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
		response = &models.SyncReply{}
	}

	return response, nil // TODO: implement
}
