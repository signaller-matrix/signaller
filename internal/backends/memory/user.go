package memory

import (
	"sync"
	"time"

	"github.com/signaller-matrix/signaller/internal"
	"github.com/signaller-matrix/signaller/internal/models"
	"github.com/signaller-matrix/signaller/internal/models/common"
	"github.com/signaller-matrix/signaller/internal/models/createroom"
	"github.com/signaller-matrix/signaller/internal/models/devices"
	"github.com/signaller-matrix/signaller/internal/models/rooms"
)

type User struct {
	name     string
	password string
	Tokens   map[string]Token
	filters  map[string]common.Filter

	backend *Backend

	mutex sync.RWMutex
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

func (user *User) CreateRoom(request createroom.Request) (internal.Room, models.ApiError) {
	for _, existingRoom := range user.backend.rooms {
		if existingRoom.AliasName() == request.RoomAliasName { // TODO: strip and check request room alias name before use
			return nil, models.NewError(models.M_ROOM_IN_USE, "")
		}
	}

	t := time.Now()

	events := make([]RoomEvent, 0)

	// Create room event
	events = append(events, RoomEvent{
		Content:        nil,
		Type:           rooms.Create,
		EventID:        internal.RandomString(eventIDSize),
		Sender:         user,
		OriginServerTS: t})

	// TODO: Add join room event

	// Set join rules event
	events = append(events, RoomEvent{
		Content:        []byte(request.Visibility), // TODO: check visibility vs join rules
		Type:           rooms.JoinRules,
		EventID:        internal.RandomString(eventIDSize),
		Sender:         user,
		OriginServerTS: t})

	// Set room name event
	if request.Name != "" {
		events = append(events, RoomEvent{
			Content:        nil, // TODO: add
			Type:           rooms.Name,
			EventID:        internal.RandomString(eventIDSize),
			Sender:         user,
			OriginServerTS: t})
	}

	// Set room alias event
	if request.RoomAliasName != "" {
		events = append(events, RoomEvent{
			Content:        nil, // TODO: add
			Type:           rooms.CanonicalAlias,
			EventID:        internal.RandomString(eventIDSize),
			Sender:         user,
			OriginServerTS: t})
	}

	room := &Room{
		id:         internal.RandomString(groupIDSize),
		aliasName:  request.RoomAliasName,
		name:       request.Name,
		topic:      request.Topic,
		events:     events,
		creator:    user,
		joined:     []internal.User{user},
		visibility: request.Visibility,
		server:     user.backend,
		state:      request.Preset}

	for i, _ := range room.events {
		room.events[i].Room = room
		//v.Room = room
	}

	user.backend.rooms[room.ID()] = room

	return room, nil
}

func (user *User) SetTopic(room internal.Room, topic string) models.ApiError {
	room.(*Room).mutex.Lock()
	defer room.(*Room).mutex.Unlock()

	if room.(*Room).creator.ID() != user.ID() { // TODO: currently only creator can change topic
		return models.NewError(models.M_FORBIDDEN, "")
	}

	room.(*Room).topic = topic
	room.(*Room).events = append(room.(*Room).events, RoomEvent{
		Type:           rooms.Topic,
		Sender:         user,
		OriginServerTS: time.Now(),
		Room:           room})

	return nil
}

func (user *User) Invite(room internal.Room, invitee internal.User) models.ApiError {
	memRoom := room.(*Room)

	memRoom.mutex.Lock()
	defer memRoom.mutex.Unlock()

	userInRoom := false

	for _, roomUser := range memRoom.joined {
		if user.ID() == roomUser.ID() {
			userInRoom = true
		}
	}

	if !userInRoom {
		return models.NewError(models.M_FORBIDDEN, "the inviter is not currently in the room") // TODO: check code
	}

	// TODO: remove repeated cycle
	for _, roomUser := range memRoom.joined {
		if roomUser.ID() == invitee.ID() {
			return models.NewError(models.M_FORBIDDEN, "the invitee is already a member of the room.") // TODO: check code
		}
	}

	for _, inviteeUser := range memRoom.invites {
		if inviteeUser.ID() == invitee.ID() {
			return models.NewError(models.M_FORBIDDEN, "user already has been invited") // TODO: check code
		}
	}

	memRoom.invites = append(memRoom.invites, invitee) // TODO: add invite event + info about inviter

	return nil
}

func (user *User) LeaveRoom(room internal.Room) models.ApiError {
	room.(*Room).mutex.Lock()
	defer room.(*Room).mutex.Unlock()

	for i, roomMember := range room.(*Room).joined {
		if roomMember.ID() == user.ID() {
			room.(*Room).joined = append(room.(*Room).joined[:i], room.(*Room).joined[i+1:]...) // TODO: add event
			return nil
		}
	}

	return models.NewError(models.M_BAD_STATE, "you are not a member of group") // TODO: check error code
}

func (user *User) SendMessage(room internal.Room, text string) models.ApiError {
	room.(*Room).mutex.Lock()
	defer room.(*Room).mutex.Unlock()

	userInRoom := false
	for _, roomMember := range room.(*Room).joined {
		if roomMember.ID() == user.ID() {
			userInRoom = true
		}
	}

	if !userInRoom {
		return models.NewError(models.M_FORBIDDEN, "")
	}

	room.(*Room).events = append(room.(*Room).events, RoomEvent{
		Content:        nil,
		Type:           rooms.Message,
		EventID:        internal.RandomString(defaultTokenSize),
		Sender:         user,
		OriginServerTS: time.Now(),
		Room:           room})

	return nil
}

func (user *User) JoinedRooms() []internal.Room {
	user.backend.mutex.Lock()
	defer user.backend.mutex.Unlock()

	var result []internal.Room

	for _, room := range user.backend.rooms {
		for _, user := range room.(*Room).joined {
			if user.ID() == user.ID() {
				result = append(result, room)
			}
		}
	}

	return result
}

func (user *User) Devices() []devices.Device {
	user.backend.mutex.Lock()
	defer user.backend.mutex.Unlock()

	var result []devices.Device

	for _, token := range user.Tokens {
		device := devices.Device{
			DeviceID: token.Device}

		result = append(result, device)
	}

	return result
}

func (user *User) SetRoomVisibility(room internal.Room, visibilityType createroom.VisibilityType) models.ApiError {
	if user.ID() != room.Creator().ID() {
		return models.NewError(models.M_FORBIDDEN, "only room owner can change visibility") // TODO: room administrators can use this method too
	}

	room.(*Room).mutex.Lock()
	defer room.(*Room).mutex.Unlock()

	room.(*Room).visibility = visibilityType

	return nil
}

func (user *User) ChangePassword(newPassword string) {
	user.mutex.Lock()
	defer user.mutex.Unlock()

	user.password = newPassword
}

func (user *User) Logout(token string) {
	delete(user.Tokens, token)
}

func (user *User) LogoutAll() {
	user.Tokens = make(map[string]Token)
}

func (user *User) JoinRoom(room internal.Room) models.ApiError {
	memRoom := room.(*Room)

	memRoom.mutex.Lock()
	defer memRoom.mutex.Unlock()

	for _, roomUser := range memRoom.joined {
		if roomUser.ID() == user.ID() {
			return models.NewError(models.M_BAD_STATE, "user already in room") // TODO: check code
		}
	}

	memRoom.joined = append(memRoom.joined, user)

	return nil
}

func (user *User) AddFilter(filterID string, filter common.Filter) {
	user.mutex.Lock()
	defer user.mutex.Unlock()

	user.filters[filterID] = filter
}

func (user *User) GetFilterByID(filterID string) *common.Filter {
	user.mutex.RLock()
	defer user.mutex.RUnlock()

	if filterReq, ok := user.filters[filterID]; ok {
		return &filterReq
	}

	return nil
}
