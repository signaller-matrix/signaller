package memory

import (
	"testing"

	"github.com/signaller-matrix/signaller/internal/models/createroom"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoom(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, err := backend.Register("user1", "", "")
	assert.NoError(t, err)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic",
		Preset:        createroom.PublicChat}

	room, err := user.CreateRoom(request)
	assert.NoError(t, err)
	assert.Equal(t, request.RoomAliasName, room.AliasName())
	assert.Equal(t, request.Name, room.Name())
	assert.Equal(t, request.Topic, room.Topic())
	assert.Equal(t, user.ID(), room.Creator().ID())
	assert.Equal(t, request.Preset, room.State())
	assert.Equal(t, 1, len(backend.rooms))
	assert.Equal(t, "!"+room.(*Room).id+":"+backend.hostname, room.ID())
}

func TestCreateAlreadyExistingRoom(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, _ := backend.Register("user1", "", "")

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	_, err := user.CreateRoom(request)
	assert.NoError(t, err)

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
	err := user.SetTopic(room, newTopic)
	assert.NoError(t, err)
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
	err := user2.SetTopic(room, newTopic)
	assert.NotNil(t, err)
}

func TestLeaveRoom(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, _ := backend.Register("user1", "", "")

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	room, _ := user.CreateRoom(request)

	assert.Equal(t, 1, len(room.(*Room).joined))

	err := user.LeaveRoom(room)
	assert.Equal(t, 0, len(room.(*Room).joined))

	// Try to leave room again must throw error
	err = user.LeaveRoom(room)
	assert.NotNil(t, err)
}

func TestRoomUserCount(t *testing.T) {
	backend := NewBackend("localhost")

	user1, _, err := backend.Register("user1", "", "")
	assert.NoError(t, err)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	room, err := user1.CreateRoom(request)
	assert.NoError(t, err)
	assert.Len(t, room.Users(), 1)

	// TODO: add join another user test
}
