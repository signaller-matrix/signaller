package sendmessage

// SendMessageReply represents reply for send message command
// https://matrix.org/docs/spec/client_server/r0.4.0.html#put-matrix-client-r0-rooms-roomid-state-eventtype-statekey
type SendMessageReply struct {
	EventID string `json:"event_id"` // A unique identifier for the event.
}
