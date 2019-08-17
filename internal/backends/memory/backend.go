package memory

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/signaller-matrix/signaller/internal"
	"github.com/signaller-matrix/signaller/internal/models"
	"github.com/signaller-matrix/signaller/internal/models/common"
	"github.com/signaller-matrix/signaller/internal/models/createroom"
	"github.com/signaller-matrix/signaller/internal/models/events"
	"github.com/wangjia184/sortedset"
)

type Backend struct {
	data                 map[string]internal.User
	rooms                map[string]internal.Room
	events               *sortedset.SortedSet
	roomAliases          map[string]internal.Room
	hostname             string
	validateUsernameFunc func(string) error // TODO: create ability to redefine validation func
	mutex                sync.RWMutex
}

type Token struct {
	Device string
}

func NewBackend(hostname string) *Backend {
	return &Backend{
		hostname:             hostname,
		validateUsernameFunc: defaultValidationUsernameFunc,
		rooms:                make(map[string]internal.Room),
		roomAliases:          make(map[string]internal.Room),
		events:               sortedset.New(),
		data:                 make(map[string]internal.User)}
}

func (backend *Backend) Register(username, password, device string) (user internal.User, token string, err models.ApiError) {
	backend.mutex.Lock()

	if backend.validateUsernameFunc != nil {
		err := backend.validateUsernameFunc(username)
		if err != nil {
			return nil, "", models.NewError(models.M_INVALID_USERNAME, err.Error())
		}
	}

	if _, ok := backend.data[username]; ok {
		backend.mutex.Unlock()
		return nil, "", models.NewError(models.M_USER_IN_USE, "trying to register a user ID which has been taken")
	}

	user = &User{
		name:     username,
		password: password,
		Tokens:   make(map[string]Token),
		backend:  backend,
		filters:  make(map[string]common.Filter)}

	backend.data[username] = user

	backend.mutex.Unlock()
	return backend.Login(username, password, device)
}

func (backend *Backend) Login(username, password, device string) (user internal.User, token string, err models.ApiError) {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	user, ok := backend.data[username]
	if !ok {
		return nil, "", models.NewError(models.M_FORBIDDEN, "wrong username")
	}

	if user.Password() != password {
		return nil, "", models.NewError(models.M_FORBIDDEN, "wrong password")
	}

	token = internal.RandomString(defaultTokenSize)

	backend.data[username].(*User).Tokens[token] = Token{Device: device}

	return user, token, nil
}

func (backend *Backend) GetUserByToken(token string) internal.User {
	backend.mutex.RLock()
	defer backend.mutex.RUnlock()

	for _, user := range backend.data {
		for userToken := range user.(*User).Tokens {
			if userToken == token {
				return user
			}
		}
	}

	return nil
}

func (backend *Backend) GetRoomByID(id string) internal.Room {
	backend.mutex.RLock()
	defer backend.mutex.RUnlock()

	for roomID, room := range backend.rooms {
		if roomID == id {
			return room
		}
	}

	return nil
}

func (backend *Backend) GetUserByName(userName string) internal.User {
	backend.mutex.RLock()
	defer backend.mutex.RUnlock()

	if user, exists := backend.data[userName]; exists {
		return user
	}

	return nil
}

func (backend *Backend) PublicRooms(filter string) []internal.Room {
	backend.mutex.RLock()
	defer backend.mutex.RUnlock()

	var rooms []internal.Room

	for _, room := range backend.rooms {
		if room.State() == createroom.PublicChat &&
			(strings.Contains(room.Name(), filter) ||
				strings.Contains(room.Topic(), filter) ||
				strings.Contains(room.AliasName(), filter)) {
			rooms = append(rooms, room)
		}
	}

	sort.Sort(BySize(rooms))

	return rooms
}

func (backend *Backend) GetRoomByAlias(alias string) internal.Room {
	backend.mutex.RLock()
	defer backend.mutex.RUnlock()

	alias = internal.StripAlias(backend.hostname, alias)

	if room, exists := backend.roomAliases[alias]; exists {
		return room
	}

	return nil
}

func (backend *Backend) ValidateUsernameFunc() func(string) error {
	backend.mutex.RLock()
	defer backend.mutex.RUnlock()

	return backend.validateUsernameFunc
}

func defaultValidationUsernameFunc(userName string) error {
	const re = `^\w{5,}$`

	if !regexp.MustCompile(re).MatchString(userName) {
		return fmt.Errorf("username does not match %s", re)
	}

	return nil
}

func (backend *Backend) GetEventByID(id string) events.Event {
	backend.mutex.RLock()
	defer backend.mutex.RUnlock()

	return backend.events.GetByKey(id).Value.(events.Event)
}

func (backend *Backend) PutEvent(event events.Event) error {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	backend.events.AddOrUpdate(event.ID(), sortedset.SCORE(time.Now().Unix()), event)

	return nil
}

func (backend *Backend) GetEventsSince(user internal.User, sinceToken string, limit int) []events.Event {
	sinceEventNode := backend.events.GetByKey(sinceToken)
	sEvents := backend.events.GetByScoreRange(sinceEventNode.Score(), -1, &sortedset.GetByScoreRangeOptions{
		Limit: limit,
	})

	eventsSlice := extractEventsFromNodes(sEvents)

	var returnEvents []events.Event
	for _, event := range eventsSlice {
		if isEventRelatedToUser(event, user) {
			returnEvents = append(returnEvents, event)
		}
	}

	return returnEvents
}

func extractEventsFromNodes(nodes []*sortedset.SortedSetNode) []events.Event {
	var eventsSlice []events.Event
	for _, e := range nodes {
		eventsSlice = append(eventsSlice, e.Value.(events.Event))
	}

	return eventsSlice
}

func isEventRelatedToUser(event events.Event, user internal.User) bool {
	if roomEvent, ok := event.(*events.RoomEvent); ok {
		if internal.InArray(roomEvent.RoomID, extractRoomIDsFromModel(user.JoinedRooms())) { // TODO check for invited or archived rooms
			return true
		}
	}

	return false
}

func extractRoomIDsFromModel(rooms []internal.Room) []string {
	var roomIDs []string
	for _, room := range rooms {
		roomIDs = append(roomIDs, room.ID())
	}

	return roomIDs
}
