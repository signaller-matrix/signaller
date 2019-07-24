package memory

import (
	"testing"

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
}

func TestRegisterUserWithAlreadyTakenName(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user1"
	)

	_, _, err := backend.Register(userName, "", "")
	assert.Nil(t, err)

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
	assert.Nil(t, err)

	token, err := backend.Login(userName, password, "")
	assert.Nil(t, err)
	assert.NotZero(t, token)
}

func TestLoginWithWrongCredentials(t *testing.T) {
	backend := NewBackend("localhost")

	var (
		userName = "user1"
		password = "password1"
	)

	_, _, err := backend.Register(userName, password, "")
	assert.Nil(t, err)

	_, err = backend.Login(userName, "wrong password", "")
	assert.NotNil(t, err)

	_, err = backend.Login("wrong user name", password, "")
	assert.NotNil(t, err)
}

func TestLogout(t *testing.T) {
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

	user.Logout(token)
}
