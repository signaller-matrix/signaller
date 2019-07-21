package createroom

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-createroom
type Reply struct {
	RoomID string `json:"room_id"`
}
