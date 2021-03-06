package internal

import (
	"github.com/signaller-matrix/signaller/internal/models"
	"github.com/signaller-matrix/signaller/internal/models/common"
	"github.com/signaller-matrix/signaller/internal/models/createroom"
	"github.com/signaller-matrix/signaller/internal/models/devices"
	"github.com/signaller-matrix/signaller/internal/models/events"
	"github.com/signaller-matrix/signaller/internal/models/sync"
)

type Backend interface {
	Register(username, password, device string) (user User, token string, err models.ApiError)
	Login(username, password, device string) (user User, token string, err models.ApiError)
	GetUserByToken(token string) (user User)
	GetUserByName(userName string) User
	GetRoomByID(id string) Room
	PublicRooms(filter string) []Room
	ValidateUsernameFunc() func(string) error
	GetEventByID(id string) events.Event
	PutEvent(events.Event) error
	GetRoomByAlias(string) Room
	GetEventsSince(user User, sinceToken string, limit int) []events.Event
}

type Room interface {
	ID() string
	Creator() User
	Users() []User
	AliasName() string
	Aliases() []string
	Name() string
	Topic() string
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
	CreateRoom(request createroom.Request) (Room, models.ApiError)
	LeaveRoom(room Room) models.ApiError
	SetTopic(room Room, topic string) models.ApiError
	SendMessage(room Room, text string) models.ApiError
	JoinedRooms() []Room
	ChangePassword(newPassword string)
	Devices() []devices.Device
	SetRoomVisibility(Room, createroom.VisibilityType) models.ApiError
	Logout(token string)
	LogoutAll()
	JoinRoom(Room) models.ApiError
	Invite(Room, User) models.ApiError
	AddFilter(filterID string, filter common.Filter)
	GetFilterByID(filterID string) *common.Filter
	AddRoomAlias(Room, string) models.ApiError
	DeleteRoomAlias(string) models.ApiError
	Sync(token string, request sync.SyncRequest) (response *sync.SyncReply, err models.ApiError)
}
