package memory

import (
	"sync"

	"github.com/nxshock/signaller/internal"
	"github.com/nxshock/signaller/internal/models/createroom"
	"github.com/nxshock/signaller/internal/models/rooms"
)

type Room struct {
	id         string
	Visibility createroom.VisibilityType
	aliasName  string
	name       string
	topic      string

	creator internal.User

	events []rooms.Event

	mutex sync.RWMutex
}

func (room *Room) ID() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.id
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

func (room *Room) Users() internal.User {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return nil // TODO: implement
}

func (room *Room) Events() []rooms.Event {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.events
}

func (room *Room) Creator() internal.User {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.creator
}

func (room *Room) NewEvent(event rooms.Event) {
	room.mutex.Lock()
	defer room.mutex.Unlock()

	room.events = append(room.events, event)
}
