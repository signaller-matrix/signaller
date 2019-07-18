package matrix

type JoinedRoomReply struct {
	JoinedRooms []string `json:"joined_rooms"`
}

// SendMessageReply represents reply for send message command
// https://matrix.org/docs/spec/client_server/r0.4.0.html#put-matrix-client-r0-rooms-roomid-state-eventtype-statekey
type SendMessageReply struct {
	EventID string `json:"event_id"` // A unique identifier for the event.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-createroom
type CreateRoomReply struct {
	RoomID string `json:"room_id"` // Information about the newly created room.
}

type GetLoginReply struct {
	Flows []Flow // The homeserver's supported login types
}

type LoginReply struct {
	AccessToken string `json:"access_token"`
	HomeServer  string `json:"home_server,omitempty"` // TODO: check api
	UserID      string `json:"user_id"`
}

type SyncReply struct {
	NextBatch              string         `json:"next_batch"`                 // Required. The batch token to supply in the since param of the next /sync request.
	Rooms                  RoomsSyncReply `json:"rooms"`                      // Updates to rooms.
	Presence               Presence       `json:"presence"`                   // The updates to the presence status of other users.
	AccountData            AccountData    `json:"account_data"`               // The global private data created by this user.
	ToDevice               ToDevice       `json:"to_device"`                  // Information on the send-to-device messages for the client device, as defined in Send-to-Device messaging.
	DeviceLists            DeviceLists    `json:"device_lists"`               // Information on end-to-end device updates, as specified in End-to-end encryption.
	DeviceOneTimeKeysCount map[string]int `json:"device_one_time_keys_count"` // Information on end-to-end encryption keys, as specified in End-to-end encryption.
}

type RoomsSyncReply struct {
	Join   map[string]JoinedRoom  `json:"join"`   // The rooms that the user has joined.
	Invite map[string]InvitedRoom `json:"invite"` // The rooms that the user has been invited to.
	Leave  map[string]LeftRoom    `json:"leave"`  // The rooms that the user has left or been banned from.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id276
type JoinRoomReply struct {
	RoomID string `json:"room_id"` // The joined room ID must be returned in the room_id field.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-versions
type VersionsReply struct {
	Versions         []string        `json:"versions"` // The supported versions.
	UnstableFeatures map[string]bool `json:"unstable_features,omitempty"`
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-account-whoami
type WhoAmIReply struct {
	UserID string `json:"user_id"` // Required. The user id that owns the access token.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-rooms-roomid-members
type MembersReply struct {
	Chunk []MemberEvent `json:"chunk"`
}

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-register
type RegisterResponse struct {
	UserID      string `json:"user_id"`                // Required. The fully-qualified Matrix user ID (MXID) that has been registered. Any user ID returned by this API must conform to the grammar given in the Matrix specification.
	AccessToken string `json:"access_token,omitempty"` // An access token for the account. This access token can then be used to authorize other requests. Required if the inhibit_login option is false.
	DeviceID    string `json:"device_id,omitempty"`    // ID of the registered device. Will be the same as the corresponding parameter in the request, if one was specified. Required if the inhibit_login option is false.
}
