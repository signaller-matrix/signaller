package memory

import (
	"fmt"
	"sync"
	"time"

	"github.com/signaller-matrix/signaller/internal"
	"github.com/signaller-matrix/signaller/internal/models"
	"github.com/signaller-matrix/signaller/internal/models/common"
	"github.com/signaller-matrix/signaller/internal/models/createroom"
	"github.com/signaller-matrix/signaller/internal/models/devices"
	"github.com/signaller-matrix/signaller/internal/models/events"
	mSync "github.com/signaller-matrix/signaller/internal/models/sync"
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

	currentUnixTime := time.Now().Unix()

	eventsSlice := make([]events.RoomEvent, 0)

	// Create room event
	eventsSlice = append(eventsSlice, events.RoomEvent{
		ContentData:    nil,
		EType:          events.Create,
		EventID:        internal.RandomString(eventIDSize),
		Sender:         user.ID(),
		OriginServerTs: currentUnixTime})

	// TODO: Add join room event

	// Set join rules event
	eventsSlice = append(eventsSlice, events.RoomEvent{
		ContentData:    []byte(request.Visibility), // TODO: check visibility vs join rules
		EType:          events.JoinRules,
		EventID:        internal.RandomString(eventIDSize),
		Sender:         user.ID(),
		OriginServerTs: currentUnixTime})

	// Set room name event
	if request.Name != "" {
		eventsSlice = append(eventsSlice, events.RoomEvent{
			ContentData:    nil, // TODO: add
			EType:          events.Name,
			EventID:        internal.RandomString(eventIDSize),
			Sender:         user.ID(),
			OriginServerTs: currentUnixTime})
	}

	// Set room alias event
	if request.RoomAliasName != "" {
		eventsSlice = append(eventsSlice, events.RoomEvent{
			ContentData:    nil, // TODO: add
			EType:          events.CanonicalAlias,
			EventID:        internal.RandomString(eventIDSize),
			Sender:         user.ID(),
			OriginServerTs: currentUnixTime})
	}

	room := &Room{
		id:         internal.RandomString(groupIDSize),
		aliasName:  request.RoomAliasName,
		name:       request.Name,
		topic:      request.Topic,
		creator:    user,
		joined:     []internal.User{user},
		visibility: request.Visibility,
		server:     user.backend,
		state:      request.Preset}

	for i, _ := range eventsSlice {
		eventsSlice[i].RoomID = room.id
	}

	for i, _ := range eventsSlice {
		user.backend.PutEvent(&eventsSlice[i])
	}

	user.backend.rooms[room.ID()] = room

	return room, nil
}

func (user *User) SetTopic(room internal.Room, topic string) models.ApiError {
	memRoom := room.(*Room)

	memRoom.mutex.Lock()

	if memRoom.creator.ID() != user.ID() { // TODO: currently only creator can change topic
		memRoom.mutex.Unlock()
		return models.NewError(models.M_FORBIDDEN, "")
	}

	memRoom.topic = topic

	memRoom.mutex.Unlock()

	rEvent := &events.RoomEvent{
		EType:          events.Topic,
		Sender:         user.ID(),
		OriginServerTs: time.Now().Unix(),
		RoomID:         memRoom.id}

	user.backend.PutEvent(rEvent)

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
	memRoom := room.(*Room)

	memRoom.mutex.Lock()
	defer memRoom.mutex.Unlock()

	for i, roomMember := range room.(*Room).joined {
		if roomMember.ID() == user.ID() {
			room.(*Room).joined = append(room.(*Room).joined[:i], room.(*Room).joined[i+1:]...) // TODO: add event
			return nil
		}
	}

	return models.NewError(models.M_BAD_STATE, "you are not a member of group") // TODO: check error code
}

func (user *User) SendMessage(room internal.Room, text string) models.ApiError {
	memRoom := room.(*Room)

	memRoom.mutex.RLock()
	defer memRoom.mutex.RUnlock()

	userInRoom := false
	for _, roomMember := range memRoom.joined {
		if roomMember.ID() == user.ID() {
			userInRoom = true
		}
	}

	if !userInRoom {
		return models.NewError(models.M_FORBIDDEN, "")
	}

	rEvent := &events.RoomEvent{
		ContentData:    nil,
		EType:          events.Message,
		EventID:        internal.RandomString(defaultTokenSize),
		Sender:         user.ID(),
		OriginServerTs: time.Now().Unix(),
		RoomID:         memRoom.id}

	user.backend.PutEvent(rEvent)

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

	memRoom := room.(*Room)

	memRoom.mutex.Lock()
	defer memRoom.mutex.Unlock()

	memRoom.visibility = visibilityType

	return nil
}

func (user *User) ChangePassword(newPassword string) {
	user.mutex.Lock()
	defer user.mutex.Unlock()

	user.password = newPassword
}

func (user *User) Logout(token string) {
	user.mutex.Lock()
	defer user.mutex.Unlock()

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

func (user *User) AddRoomAlias(room internal.Room, alias string) models.ApiError {
	user.backend.mutex.Lock()
	defer user.backend.mutex.Unlock()

	if room.Creator().ID() != user.ID() {
		return models.NewError(models.M_FORBIDDEN, "only room creator can add room alias") // TODO: make room admins can use this method
	}

	if _, exists := user.backend.roomAliases[alias]; exists {
		return models.NewError(models.M_UNKNOWN, fmt.Sprintf("room alias #%s:%s already exists", alias, user.backend.hostname))
	}

	user.backend.roomAliases[alias] = room

	return nil
}

func (user *User) DeleteRoomAlias(alias string) models.ApiError {
	user.backend.mutex.Lock()
	defer user.backend.mutex.Unlock()

	room := user.backend.GetRoomByAlias(alias)
	if room == nil {
		return models.NewError(models.M_NOT_FOUND, "room not found")
	}

	if room.Creator().ID() != user.ID() {
		return models.NewError(models.M_FORBIDDEN, "only room creator can delete room alias") // TODO: make room admins can use this method
	}

	delete(user.backend.roomAliases, alias)

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

func (user *User) Sync(token string, request mSync.SyncRequest) (response *mSync.SyncReply, err models.ApiError) {
	response = mSync.BuildEmptySyncReply()
	emptyEventList := true
	emptySinceToken := true
	var eventsList []events.Event

	if request.Since != "" {
		emptySinceToken = false
		eventsList := user.backend.GetEventsSince(user, request.Since, 0) // TODO filtering

		if len(eventsList) != 0 {
			emptyEventList = false
		}
	}

	if !emptyEventList {
		for _, room := range user.JoinedRooms() {
			filteredEventList := filterEventsByRoom(room.ID(), eventsList)
			var prevBatch string
			if len(filteredEventList) != 0 {
				prevBatch = filteredEventList[0].ID()
			}
			response.Rooms.Join[room.ID()] = mSync.JoinedRoom{
				RoomSummary: mSync.RoomSummary{
					Heroes:             nil,
					JoinedMemberCount:  0,
					InvitedMemberCount: 0,
				},
				State: events.State{
					Events: nil,
				},
				Timeline: mSync.Timeline{
					Events:    filteredEventList,
					Limited:   false, // TODO filtering
					PrevBatch: prevBatch,
				},
			}
		}
		response.NextBatch = eventsList[len(eventsList)-1].ID()
	} else if emptySinceToken {
		// TODO
	} else {
		// TODO wait for new events or just return empty response when timeout is reached
	}

	return response, nil
}

func filterEventsByRoom(roomID string, eventList []events.Event) []events.RoomEvent {
	var filteredEventList []events.RoomEvent
	for _, event := range eventList {
		if roomEvent, ok := event.(*events.RoomEvent); ok {
			if roomEvent.RoomID == roomID {
				filteredEventList = append(filteredEventList, *event.(*events.RoomEvent))
			}
		}
	}

	return filteredEventList
}
