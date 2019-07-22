package memory

import (
	"crypto/rand"
	"fmt"
)

// newToken returns new generated token with specified length
func newToken(size int) string {
	b := make([]byte, size)
	rand.Read(b) // TODO: check may be can be error

	return fmt.Sprintf("%x", b)
}
