package internal

import (
	"github.com/nxshock/signaller/internal/models"
	"github.com/nxshock/signaller/internal/models/createroom"
	"github.com/nxshock/signaller/internal/models/rooms"
	"github.com/nxshock/signaller/internal/models/sync"
)

type Backend interface {
	Register(username, password, device string) (user User, token string, err *models.ApiError)
	Login(username, password, device string) (token string, err *models.ApiError)
	Logout(token string) *models.ApiError
	Sync(token string, request sync.SyncRequest) (response *sync.SyncReply, err *models.ApiError)
}

type Room interface {
	ID() string
	Creator() User
	Users() User
	AliasName() string
	Name() string
	Topic() string
	SetTopic(user User, topic string) *models.ApiError
	Events() []rooms.Event
}

type User interface {
	Name() string
	ID() string
	Password() string
	CreateRoom(request createroom.Request) (Room, *models.ApiError)
}
