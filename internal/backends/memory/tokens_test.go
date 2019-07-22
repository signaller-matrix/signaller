package memory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenGenerator(t *testing.T) {
	token := newToken(defaultTokenSize)
	assert.Len(t, token, defaultTokenSize*2)
}
