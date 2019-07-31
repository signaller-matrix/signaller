package memory

import (
	"testing"

	"github.com/nxshock/signaller/internal"

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

func TestGetUserByToken(t *testing.T) {
	backend := NewBackend("localhost")

	user, token, err := backend.Register("user1", "", "")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	gotUser := backend.GetUserByToken(token)
	assert.Equal(t, user, gotUser)
}

func TestGetUserByWrongToken(t *testing.T) {
	backend := NewBackend("localhost")

	_, token, err := backend.Register("user1", "", "")
	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	gotUser := backend.GetUserByToken("wrong token")
	assert.Nil(t, gotUser)
}

func TestLogoutWithWrongToken(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user1"
		password = "password1"
	)

	user, _, err := backend.Register(userName, password, "")
	assert.Nil(t, err)

	token, err := backend.Login(userName, password, "")
	assert.Nil(t, err)
	assert.NotZero(t, token)

	user.Logout("worng token")
}

func TestJoinedRooms(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, err := backend.Register("user1", "", "")
	assert.Nil(t, err)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Topic:         "topic"}

	room, err := user.CreateRoom(request)
	assert.Nil(t, err)

	rooms := user.JoinedRooms()
	assert.Equal(t, []internal.Room{room}, rooms)
}

func TestNewPassword(t *testing.T) {
	backend := NewBackend("localhost")

	var newPassword = "new password"

	user, _, err := backend.Register("user1", "old password", "")
	assert.Nil(t, err)

	user.ChangePassword(newPassword)
	assert.Equal(t, newPassword, user.Password())
}

func TestDevices(t *testing.T) {
	backend := NewBackend("localhost")

	var expectedDeviceID = "my device"

	user, _, err := backend.Register("user1", "", expectedDeviceID)
	assert.Nil(t, err)

	devices := user.Devices()
	assert.Len(t, devices, 1)
	assert.Equal(t, expectedDeviceID, devices[0].DeviceID)
}

func TestSetRoomVisibility(t *testing.T) {
	backend := NewBackend("localhost")

	user, _, err := backend.Register("user1", "", "")
	assert.Nil(t, err)

	request := createroom.Request{
		RoomAliasName: "room1",
		Name:          "room1",
		Visibility:    createroom.VisibilityTypePrivate}

	room, err := user.CreateRoom(request)
	assert.Nil(t, err)
	assert.NotNil(t, room)
	assert.Equal(t, createroom.VisibilityTypePrivate, room.Visibility())

	err = user.SetRoomVisibility(room, createroom.VisibilityTypePublic)
	assert.Nil(t, err)
	assert.Equal(t, createroom.VisibilityTypePublic, room.Visibility())
}
