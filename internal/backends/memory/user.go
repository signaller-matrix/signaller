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

	t := time.Now()

	events := make([]RoomEvent, 0)

	// Create room event
	events = append(events, RoomEvent{
		Content:        nil,
		Type:           rooms.Create,
		EventID:        internal.NewToken(eventIDSize),
		Sender:         user,
		OriginServerTS: t})

	// Set join rules event
	events = append(events, RoomEvent{
		Content:        []byte(request.Visibility), // TODO: check visibility vs join rules
		Type:           rooms.JoinRules,
		EventID:        internal.NewToken(eventIDSize),
		Sender:         user,
		OriginServerTS: t})

	// Set room name event
	if request.Name != "" {
		events = append(events, RoomEvent{
			Content:        nil, // TODO: add
			Type:           rooms.Name,
			EventID:        internal.NewToken(eventIDSize),
			Sender:         user,
			OriginServerTS: t})
	}

	// Set room alias event
	if request.RoomAliasName != "" {
		events = append(events, RoomEvent{
			Content:        nil, // TODO: add
			Type:           rooms.CanonicalAlias,
			EventID:        internal.NewToken(eventIDSize),
			Sender:         user,
			OriginServerTS: t})
	}

	room := &Room{
		id:        internal.NewToken(groupIDSize),
		aliasName: request.RoomAliasName,
		name:      request.Name,
		topic:     request.Topic,
		events:    events,
		creator:   user}

	for i, _ := range room.events {
		room.events[i].Room = room
		//v.Room = room
	}

	user.backend.rooms[room.id] = room

	return room, nil
}

func (user *User) SetTopic(room internal.Room, topic string) *models.ApiError {
	room.(*Room).mutex.Lock()
	defer room.(*Room).mutex.Unlock()

	if room.(*Room).creator.ID() != user.ID() { // TODO: currently only creator can change topic
		return internal.NewError(models.M_FORBIDDEN, "")
	}

	room.(*Room).topic = topic
	room.(*Room).events = append(room.(*Room).events, RoomEvent{
		Type:           rooms.Topic,
		Sender:         user,
		OriginServerTS: time.Now(),
		Room:           room})

	return nil
}
