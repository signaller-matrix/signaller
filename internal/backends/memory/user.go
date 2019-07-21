package memory

import (
	"time"

	"github.com/nxshock/signaller/internal"
	"github.com/nxshock/signaller/internal/models"
	"github.com/nxshock/signaller/internal/models/createroom"
	"github.com/nxshock/signaller/internal/models/rooms"
)

type User struct {
	name     string
	password string
	Tokens   map[string]Token

	backend *Backend
}

func (user *User) ID() string {
	return "@" + user.name + ":" + user.backend.hostname
}

func (user *User) Name() string {
	return user.name
}

func (user *User) Password() string {
	return user.password
}

func (user *User) CreateRoom(request createroom.Request) (internal.Room, *models.ApiError) {
	for _, existingRoom := range user.backend.rooms {
		if existingRoom.AliasName() == request.RoomAliasName { // TODO: strip and check request room alias name before use
			return nil, internal.NewError(models.M_ROOM_IN_USE, "")
		}
	}

	room := &Room{
		id:        internal.NewToken(groupIDSize),
		aliasName: request.RoomAliasName,
		name:      request.Name,
		topic:     request.Topic,
		events: []RoomEvent{
			RoomEvent{
				Content:        nil,
				Type:           rooms.Create,
				EventID:        internal.NewToken(eventIDSize),
				Sender:         user,
				OriginServerTS: time.Now()}},
		creator: user}

	room.events[0].Room = room

	user.backend.rooms[room.id] = room

	return room, nil
}
