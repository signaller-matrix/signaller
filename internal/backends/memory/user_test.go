package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nxshock/signaller/internal/models/createroom"
)

func TestUserID(t *testing.T) {
	var (
		userName       = "user1"
		hostName       = "localhost"
		expectedUserID = "@user1:localhost"
	)

	backend := NewBackend(hostName)
	user, _, err := backend.Register(userName, "", "")
	assert.Nil(t, err)

	assert.Equal(t, expectedUserID, user.ID())
}

func TestUserMessage(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, err := backend.Register("user1", "", "")
	assert.Nil(t, err)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1"}

	room, err := user.CreateRoom(request)
	assert.Nil(t, err)

	err = user.SendMessage(room, "hello")
	assert.Nil(t, err)
}

func TestUserMessageInWrongRoom(t *testing.T) {
	backend := NewBackend("localhost")

	user1, _, err := backend.Register("user1", "", "")
	assert.Nil(t, err)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1"}

	room, err := user1.CreateRoom(request)
	assert.Nil(t, err)

	user2, _, err := backend.Register("user2", "", "")
	assert.Nil(t, err)

	err = user2.SendMessage(room, "hello")
	assert.NotNil(t, err)
}
