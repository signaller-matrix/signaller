package registeravailable

type Request struct {
	Username string `json:"username"` // The username to check the availability of.
}

type Response struct {
	Available bool `json:"available"` // A flag to indicate that the username is available. This should always be true when the server replies with 200 OK.
}
