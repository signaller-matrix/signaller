package memory

import (
	"sort"
	"sync"

	"github.com/nxshock/signaller/internal/models/createroom"

	"github.com/nxshock/signaller/internal"
	"github.com/nxshock/signaller/internal/models"
	mSync "github.com/nxshock/signaller/internal/models/sync"
)

type Backend struct {
	data     map[string]internal.User
	rooms    map[string]internal.Room
	hostname string
	mutex    sync.Mutex // TODO: replace with RW mutex
}

type Token struct {
	Device string
}

func NewBackend(hostname string) *Backend {
	return &Backend{
		hostname: hostname,
		rooms:    make(map[string]internal.Room),
		data:     make(map[string]internal.User)}
}

func (backend *Backend) Register(username, password, device string) (user internal.User, token string, err *models.ApiError) {
	backend.mutex.Lock()

	if _, ok := backend.data[username]; ok {
		backend.mutex.Unlock()
		return nil, "", internal.NewError(models.M_USER_IN_USE, "trying to register a user ID which has been taken")
	}

	user = &User{
		name:     username,
		password: password,
		Tokens:   make(map[string]Token),
		backend:  backend}

	backend.data[username] = user

	backend.mutex.Unlock()
	return backend.Login(username, password, device)
}

func (backend *Backend) Login(username, password, device string) (user internal.User, token string, err *models.ApiError) {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	user, ok := backend.data[username]
	if !ok {
		return nil, "", internal.NewError(models.M_FORBIDDEN, "wrong username")
	}

	if user.Password() != password {
		return nil, "", internal.NewError(models.M_FORBIDDEN, "wrong password")
	}

	token = newToken(defaultTokenSize)

	backend.data[username].(*User).Tokens[token] = Token{Device: device}

	return user, token, nil
}

func (backend *Backend) Sync(token string, request mSync.SyncRequest) (response *mSync.SyncReply, err *models.ApiError) {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	return nil, nil // TODO: implement
}

func (backend *Backend) GetUserByToken(token string) internal.User {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

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
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	for roomID, room := range backend.rooms {
		if roomID == id {
			return room
		}
	}

	return nil
}

func (backend *Backend) GetUserByName(userName string) internal.User {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	if user, exists := backend.data[userName]; exists {
		return user
	}

	return nil
}

func (backend *Backend) PublicRooms() []internal.Room {
	backend.mutex.Lock()
	defer backend.mutex.Unlock()

	var rooms []internal.Room

	for _, room := range backend.rooms {
		if room.State() == createroom.PublicChat {
			rooms = append(rooms, room)
		}
	}

	sort.Sort(BySize(rooms))

	return rooms
}
