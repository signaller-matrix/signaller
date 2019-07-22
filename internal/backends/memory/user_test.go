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
