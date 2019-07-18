package matrix

type MRelatesTo struct {
	InReplyTo MInReplyTo `json:"m.in_reply_to"`
}

type MInReplyTo struct {
	EventID string `json:"event_id"`
}

// Invite3pid represents third party IDs to invite into the room
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-createroom
type Invite3pid struct {
	IDServer string `json:"id_server"` // Required. The hostname+port of the identity server which should be used for third party identifier lookups.
	Medium   string `json:"medium"`    // Required. The kind of address being passed in the address field, for example email.
	Address  string `json:"address"`   // Required. The invitee's third party identifier.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-rooms-roomid-members
type MemberEvent struct {
	Content        EventContent `json:"content"`          // Required.
	Type           string       `json:"type"`             //Required. Must be 'm.room.member'.
	EventID        string       `json:"event_id"`         // Required. The globally unique event identifier.
	Sender         string       `json:"sender"`           // Required. Contains the fully-qualified ID of the user who sent this event.
	OriginServerTs int          `json:"origin_server_ts"` // Required. Timestamp in milliseconds on originating homeserver when this event was sent.
	Unsigned       UnsignedData `json:"unsigned"`         // Contains optional extra information about the event.
	RoomID         string       `json:"room_id"`          // Required. The ID of the room associated with this event. Will not be present on events that arrive through /sync, despite being required everywhere else.
	PrevContent    EventContent `json:"prev_content"`     // Optional. The previous content for this event. If there is no previous content, this key will be missing.
	StateKey       string       `json:"state_key"`        // Required. The user_id this membership event relates to. In all cases except for when membership is join, the user ID sending the event does not need to match the user ID in the state_key, unlike other events. Regular authorisation rules still apply.
}
