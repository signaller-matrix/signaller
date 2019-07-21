package memory

import (
	"strings"
	"testing"

	"github.com/nxshock/signaller/internal/models/createroom"

	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		username = "user1"
		password = "password1"
		device   = "device1"
	)

	user, token, err := backend.Register(username, password, device)
	assert.Nil(t, err)
	assert.Equal(t, username, user.Name())
	assert.Equal(t, password, user.Password())
	assert.NotEmpty(t, token)
	assert.True(t, strings.HasSuffix(user.ID(), backend.hostname))
}

func TestCreateRoom(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, err := backend.Register("user1", "", "")
	assert.Nil(t, err)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	room, err := user.CreateRoom(request)
	assert.Nil(t, err)
	assert.Equal(t, request.RoomAliasName, room.AliasName())
	assert.Equal(t, request.Name, room.Name())
	assert.Equal(t, request.Topic, room.Topic())
	assert.Equal(t, user.ID(), room.Creator().ID())
	assert.Equal(t, 1, len(backend.rooms))
}

func TestCreateAlreadyExistingRoom(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, _ := backend.Register("user1", "", "")

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	_, err := user.CreateRoom(request)
	assert.Nil(t, err)

	_, err = user.CreateRoom(request)
	assert.NotNil(t, err)
}

func TestSetRoomTopic(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, _ := backend.Register("user1", "", "")

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	room, _ := user.CreateRoom(request)

	var newTopic = "new topic"
	err := room.SetTopic(user, newTopic)
	assert.Nil(t, err)
	assert.Equal(t, newTopic, room.Topic())
	assert.Equal(t, 5, len(room.Events())) // TODO: check start event count
}

func TestSetRoomTopicWithnprivelegedUser(t *testing.T) {
	backend := NewBackend("localhost")

	creator, _, _ := backend.Register("user1", "", "")
	user2, _, _ := backend.Register("user2", "", "")

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	room, _ := creator.CreateRoom(request)

	var newTopic = "new topic"
	err := room.SetTopic(user2, newTopic)
	assert.NotNil(t, err)
}
