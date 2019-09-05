package memory

import (
	"sync"

	"github.com/signaller-matrix/signaller/internal/models/createroom"
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

	creator string
	joined  []string
	invites []string

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

func (room *Room) Aliases() []string {
	room.server.mutex.RLock()
	defer room.server.mutex.RUnlock()

	var aliases []string
	for alias, serverRoom := range room.server.roomAliases {
		if serverRoom.ID() == room.ID() {
			aliases = append(aliases, alias)
		}
	}

	return aliases
}

func (room *Room) Topic() string {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.topic
}

func (this *Room) Users() []string {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	return this.joined
}

func (room *Room) Visibility() createroom.VisibilityType {
	room.mutex.RLock()
	defer room.mutex.RUnlock()

	return room.visibility
}

func (this *Room) Creator() string {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	return this.creator
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

func (this *Room) PutInvited(invitedID string) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	this.invites = append(this.invites, invitedID)
}

func (this *Room) PutJoined(joinedID string) {
	this.mutex.RLock()
	defer this.mutex.RUnlock()

	this.joined = append(this.joined, joinedID)
}
