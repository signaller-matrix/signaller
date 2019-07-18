package models

import (
	"encoding/json"
)

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id242
type SyncRequest struct {
	Filter      string      `url:"filter,omitempty"`       // The ID of a filter created using the filter API or a filter JSON object encoded as a string. The server will detect whether it is an ID or a JSON object by whether the first character is a "{" open brace. Passing the JSON inline is best suited to one off requests. Creating a filter using the filter API is recommended for clients that reuse the same filter multiple times, for example in long poll requests.
	Since       string      `url:"since,omitempty"`        // A point in time to continue a sync from.
	FullState   bool        `url:"full_state,omitempty"`   // Controls whether to include the full state for all rooms the user is a member of.
	SetPresence SetPresence `url:"set_presence,omitempty"` // Controls whether the client is automatically marked as online by polling this API. If this parameter is omitted then the client is automatically marked as online when it uses this API. Otherwise if the parameter is set to "offline" then the client is not marked as being online when it uses this API. When set to "unavailable", the client is marked as being idle. One of: ["offline", "online", "unavailable"]
	Timeout     int         `url:"timeout,omitempty"`      // The maximum time to wait, in milliseconds, before returning this request. If no events (or other data) become available before this time elapses, the server will return a response with empty fields.
}

type JoinedRoom struct {
	State               State                    `json:"state"`                // Updates to the state, between the time indicated by the since parameter, and the start of the timeline (or all state up to the start of the timeline, if since is not given, or full_state is true).
	Timeline            Timeline                 `json:"timeline"`             // The timeline of messages and state changes in the room.
	Ephemeral           Ephemeral                `json:"ephemeral"`            // The ephemeral events in the room that aren't recorded in the timeline or state of the room. e.g. typing.
	AccountData         AccountData              `json:"account_data"`         // The private data that this user has attached to this room.
	UnreadNotifications UnreadNotificationCounts `json:"unread_notifications"` // Counts of unread notifications for this room
}

type Ephemeral struct {
	Events []Event `json:"events"` // List of events.
}

type UnreadNotificationCounts struct {
	HighlightCount    int `json:"highlight_count"`    // The number of unread notifications for this room with the highlight flag set
	NotificationCount int `json:"notification_count"` // The total number of unread notifications for this room
}

type InvitedRoom struct {
	InviteState InviteState `json:"invite_state"` // The state of a room that the user has been invited to. These state events may only have the sender, type, state_key and content keys present. These events do not replace any state that the client already has for the room, for example if the client has archived the room. Instead the client should keep two separate copies of the state: the one from the invite_state and one from the archived state. If the client joins the room then the current state will be given as a delta against the archived state not the invite_state.
}

type InviteState struct {
	Events []StrippedState `json:"events"` // The StrippedState events that form the invite state.
}

type StrippedState struct {
	// TODO: в документации EventContent, хотя вроде сервер выдаёт json.RawMessage
	Content  json.RawMessage `json:"content"`   // Required. The content for the event.
	StateKey string          `json:"state_key"` // Required. The state_key for the event.
	Type     string          `json:"type"`      // Required. The type for the event.
	Sender   string          `json:"sender"`    // Required. The sender for the event.
}

type LeftRoom struct {
	State       State       `json:"state"`        // The state updates for the room up to the start of the timeline.
	Timeline    Timeline    `json:"timeline"`     // The timeline of messages and state changes in the room up to the point when the user left.
	AccountData AccountData `json:"account_data"` // The private data that this user has attached to this room.
}

type State struct {
	events []StateEvent `json:"events"` // List of events.
}

type StateEvent struct {
	// TODO: object?
	Content        json.RawMessage `json:"content"`          // Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Type           string          `json:"type"`             // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	EventID        string          `json:"event_id"`         // Required. The globally unique event identifier.
	Sender         string          `json:"sender"`           // Required. Contains the fully-qualified ID of the user who sent this event.
	OriginServerTs int             `json:"origin_server_ts"` // Required. Timestamp in milliseconds on originating homeserver when this event was sent.
	Unsigned       UnsignedData    `json:"unsigned"`         // Contains optional extra information about the event.
	PrevContent    EventContent    `json:"prev_content"`     // Optional. The previous content for this event. If there is no previous content, this key will be missing.
	StateKey       string          `json:"state_key"`        // Required. A unique key which defines the overwriting semantics for this piece of room state. This value is often a zero-length string. The presence of this key makes this event a State Event. State keys starting with an @ are reserved for referencing user IDs, such as room members. With the exception of a few events, state events set with a given user's ID as the state key MUST only be set by that user.
}

type Timeline struct {
	Events    []RoomEvent `json:"events"`     // List of events.
	Limited   bool        `json:"limited"`    // True if the number of events returned was limited by the limit on the filter.
	PrevBatch string      `json:"prev_batch"` // A token that can be supplied to the from parameter of the rooms/{roomId}/messages endpoint.
}

type RoomEvent struct {
	// TODO: object
	Content        json.RawMessage `json:"content"`          // Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Type           string          `json:"type"`             // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	EventID        string          `json:"event_id"`         // Required. The globally unique event identifier.
	Sender         string          `json:"sender"`           // Required. Contains the fully-qualified ID of the user who sent this event.
	OriginServerTs int64           `json:"origin_server_ts"` // Required. Timestamp in milliseconds on originating homeserver when this event was sent.
	Unsigned       UnsignedData    `json:"unsigned"`         // Contains optional extra information about the event.
}

type UnsignedData struct {
	Age             int    `json:"age"`              // The time in milliseconds that has elapsed since the event was sent. This field is generated by the local homeserver, and may be incorrect if the local time on at least one of the two servers is out of sync, which can cause the age to either be negative or greater than it actually is.
	RedactedBecause Event  `json:"redacted_because"` // Optional. The event that redacted this event, if any.
	TransactionID   string `json:"transaction_id"`   // The client-supplied transaction ID, if the client being given the event is the same one which sent it.
}

type Presence struct {
	events []Event `json:"events"` // List of events.
}

type AccountData struct {
	Events []Event `json:"events"` // List of events.
}

type Event struct {
	// TODO: object
	Content json.RawMessage `json:"content"` // Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Type    string          `json:"type"`    // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
}

type EventContent struct {
	AvatarURL string `json:"avatar_url"` // The avatar URL for this user, if any. This is added by the homeserver.
	// TODO: string or null
	DisplayName      string       `json:"displayname"`        // The display name for this user, if any. This is added by the homeserver.
	Membership       Membership   `json:"membership"`         // Required. The membership state of the user. One of: ["invite", "join", "knock", "leave", "ban"]
	IsDirect         bool         `json:"is_direct"`          // Flag indicating if the room containing this event was created with the intention of being a direct chat. See Direct Messaging.
	ThirdPartyInvite Invite       `json:"third_party_invite"` //
	Unsigned         UnsignedData `json:"unsigned"`           // Contains optional extra information about the event.
}

type Invite struct {
	DisplayName string `json:"display_name"` // Required. A name which can be displayed to represent the user instead of their third party identifier
	Signed      signed `json:"signed"`       // Required. A block of content which has been signed, which servers can use to verify the event. Clients should ignore this.
}

type signed struct {
	Mxid string `json:"mxid"` // Required. The invited matrix user ID. Must be equal to the user_id property of the event.
	// TODO:
	// Signatures Signatures `json:"signatures"` // Required. A single signature from the verifying server, in the format specified by the Signing Events section of the server-server API.
	Token string `json:"token"` // Required. The token property of the containing third_party_invite object.
}

// TODO: проверить правильность выбора типа
type ToDevice struct {
	events []Event `json:"events` // List of send-to-device messages
}

type DeviceLists struct {
	Changed []string `json:"changed"` // List of users who have updated their device identity keys, or who now share an encrypted room with the client since the previous sync response.
	Left    []string `json:"left"`    // List of users with whom we do not share any encrypted rooms anymore since the previous sync response.
}

type Flow struct {
	Type AuthenticationType `json:"type"`
}
