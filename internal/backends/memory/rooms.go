package memory

import (
	"sync"

	"github.com/signaller-matrix/signaller/internal"
	"github.com/signaller-matrix/signaller/internal/models/createroom"
	"github.com/signaller-matrix/signaller/internal/models/rooms"
)

type Room struct {
	id            string
	visibility    createroom.VisibilityType
	aliasName     string
	name          string
	topic         string
	state         createroom.Preset
	worldReadable bool
	guestCanJoin  bool
	avatarURL     string

	creator internal.User
	joined  []internal.User
	invites []internal.User

	events []RoomEvent

	server *Backend

	mutex sync.RWMutex
}

func (room *Room) ID() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return "!" + room.id + ":" + room.server.hostname
}

func (room *Room) Name() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.name
}

func (room *Room) AliasName() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.aliasName
}

func (room *Room) Topic() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.topic
}

func (room *Room) Users() []internal.User {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.joined
}

func (room *Room) Events() []rooms.Event {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	result := make([]rooms.Event, 0)
	for _, v := range room.events {
		result = append(result, v.ToEvent())
	}

	return result
}

func (room *Room) Visibility() createroom.VisibilityType {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.visibility
}

func (room *Room) Creator() internal.User {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.creator
}

func (room *Room) State() createroom.Preset {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.state
}

func (room *Room) WorldReadable() bool {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.worldReadable
}

func (room *Room) GuestCanJoin() bool {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.guestCanJoin
}

func (room *Room) AvatarURL() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.avatarURL
}
