package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
