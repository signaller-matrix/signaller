package joinroom

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id276
type JoinRoomReply struct {
	RoomID string `json:"room_id"` // The joined room ID must be returned in the room_id field.
}
