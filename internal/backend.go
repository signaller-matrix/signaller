package internal

import (
	"github.com/nxshock/signaller/internal/models"
	"github.com/nxshock/signaller/internal/models/createroom"
	"github.com/nxshock/signaller/internal/models/devices"
	"github.com/nxshock/signaller/internal/models/rooms"
	"github.com/nxshock/signaller/internal/models/sync"
)

type Backend interface {
	Register(username, password, device string) (user User, token string, err *models.ApiError)
	Login(username, password, device string) (user User, token string, err *models.ApiError)
	GetUserByToken(token string) (user User)
	GetUserByName(userName string) User
	GetRoomByID(id string) Room
	Sync(token string, request sync.SyncRequest) (response *sync.SyncReply, err *models.ApiError)
	PublicRooms() []Room
}

type Room interface {
	ID() string
	Creator() User
	Users() []User
	AliasName() string
	Name() string
	Topic() string
	Events() []rooms.Event
	Visibility() createroom.VisibilityType
	WorldReadable() bool
	GuestCanJoin() bool
	AvatarURL() string
	State() createroom.Preset
}

type User interface {
	Name() string
	ID() string
	Password() string
	CreateRoom(request createroom.Request) (Room, *models.ApiError)
	LeaveRoom(room Room) *models.ApiError
	SetTopic(room Room, topic string) *models.ApiError
	SendMessage(room Room, text string) *models.ApiError
	JoinedRooms() []Room
	ChangePassword(newPassword string)
	Devices() []devices.Device
	SetRoomVisibility(Room, createroom.VisibilityType) *models.ApiError
	Logout(token string)
	LogoutAll()
	JoinRoom(Room) *models.ApiError
}
