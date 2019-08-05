package memory

import (
	"testing"

	"github.com/signaller-matrix/signaller/internal"

	"github.com/stretchr/testify/assert"
)

func TestTokenGenerator(t *testing.T) {
	token := internal.RandomString(defaultTokenSize)
	assert.Len(t, token, defaultTokenSize*2)
}
