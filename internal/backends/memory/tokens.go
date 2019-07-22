package memory

import (
	"crypto/rand"
	"fmt"
)

// newToken returns new generated token with specified length
func newToken(size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", b)
}
