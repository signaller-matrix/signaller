package internal

import (
	"net/http"
	"strings"
)

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
