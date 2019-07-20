package internal

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"strings"
)

// newToken returns new generated token with specified length
func NewToken(size int) string {
	b := make([]byte, size)
	rand.Read(b)

	return fmt.Sprintf("%x", b)
}

// getTokenFromResponse returns token from request.
func getTokenFromResponse(r *http.Request) string {
	const prefix = "Bearer "

	auth, ok := r.Header["Authorization"]
	if !ok {
		return ""
	}

	for _, v := range auth {
		if strings.HasPrefix(v, prefix) {
			return strings.TrimPrefix(v, prefix)
		}
	}

	return ""
}
