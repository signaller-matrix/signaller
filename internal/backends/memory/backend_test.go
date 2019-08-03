package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nxshock/signaller/internal/models/createroom"
)

func TestRegisterUser(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		username = "user1"
		password = "password1"
		device   = "device1"
	)

	user, token, err := backend.Register(username, password, device)
	assert.NoError(t, err)
	assert.Equal(t, username, user.Name())
	assert.Equal(t, password, user.Password())
	assert.NotEmpty(t, token)
}

func TestRegisterUserWithAlreadyTakenName(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user1"
	)

	_, _, err := backend.Register(userName, "", "")
	assert.NoError(t, err)

	_, _, err = backend.Register(userName, "", "")
	assert.NotNil(t, err)
}

func TestLogin(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user1"
		password = "password1"
	)

	_, _, err := backend.Register(userName, password, "")
	assert.NoError(t, err)

	_, token, err := backend.Login(userName, password, "")
	assert.NoError(t, err)
	assert.NotZero(t, token)
}

func TestLoginWithWrongCredentials(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user1"
		password = "password1"
	)

	_, _, err := backend.Register(userName, password, "")
	assert.NoError(t, err)

	_, _, err = backend.Login(userName, "wrong password", "")
	assert.NotNil(t, err)

	_, _, err = backend.Login("wrong user name", password, "")
	assert.NotNil(t, err)
}

func TestLogout(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user1"
		password = "password1"
	)

	user, _, err := backend.Register(userName, password, "")
	assert.NoError(t, err)

	_, token, err := backend.Login(userName, password, "")
	assert.NoError(t, err)
	assert.NotZero(t, token)

	user.Logout(token)

	assert.Nil(t, backend.GetUserByToken(token))
}

func TestGetRoomByID(t *testing.T) {
	backend := NewBackend("localhost")

	user, token, err := backend.Register("user", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, token)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1"}

	room, err := user.CreateRoom(request)
	assert.NoError(t, err)
	assert.NotNil(t, room)
	assert.Equal(t, room.ID(), backend.GetRoomByID(room.ID()).ID())

	// Get room with wrong id
	room = backend.GetRoomByID("worng id")
	assert.Nil(t, room)
}

func TestGetUserByName(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user"
	)

	user, token, err := backend.Register(userName, "", "")
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, token)

	t.Run("Test picking user with username", func(_ *testing.T) {
		user2 := backend.GetUserByName(userName)
		assert.Equal(t, user, user2)
	})

	t.Run("Test picking user with wrong username", func(_ *testing.T) {
		user2 := backend.GetUserByName("wrong username")
		assert.Nil(t, user2)
	})
}

func TestPublicRooms(t *testing.T) {
	backend := NewBackend("localhost")

	user1, _, err := backend.Register("user1", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, user1)

	// Create first room
	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Preset:        createroom.PublicChat}

	room1, err := user1.CreateRoom(request)
	assert.NoError(t, err)
	assert.NotNil(t, room1)

	// Create second room
	request = createroom.Request{
		RoomAliasName: "room2",
		Name:          "room2",
		Preset:        createroom.PublicChat}

	room2, err := user1.CreateRoom(request)
	assert.NoError(t, err)
	assert.NotNil(t, room2)

	// Make room2 has 2 users
	user2, _, err := backend.Register("user2", "", "")
	assert.NoError(t, err)
	assert.NotNil(t, user2)

	err = user2.JoinRoom(room2)
	assert.NoError(t, err)

	rooms := backend.PublicRooms()
	assert.Len(t, rooms, 2)
	assert.Equal(t, rooms[0], room2)
	assert.Equal(t, rooms[1], room1)
}
