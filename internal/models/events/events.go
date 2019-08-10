package events

import "encoding/json"

type Membership string

const (
	MembershipInvite Membership = "invite"
	MembershipJoin   Membership = "join"
	MembershipKnock  Membership = "knock"
	MembershipLeave  Membership = "leave"
	MembershipBan    Membership = "ban"
)

// Type is type of event
type EventType string

const (
	// https://matrix.org/docs/spec/client_server/latest#m-room-aliases
	Aliases EventType = "m.room.aliases"

	// https://matrix.org/docs/spec/client_server/latest#m-room-canonical-alias
	CanonicalAlias EventType = "m.room.canonical_alias"

	// https://matrix.org/docs/spec/client_server/latest#m-room-create
	Create EventType = "m.room.create"

	// https://matrix.org/docs/spec/client_server/latest#m-room-join-rules
	JoinRules EventType = "m.room.join_rules"

	// https://matrix.org/docs/spec/client_server/latest#m-room-member
	Member EventType = "m.room.member"

	// https://matrix.org/docs/spec/client_server/latest#m-room-power-levels
	PowerLevels EventType = "m.room.power_levels"

	// https://matrix.org/docs/spec/client_server/latest#m-room-redaction
	Redaction EventType = "m.room.redaction"

	// https://matrix.org/docs/spec/client_server/latest#m-room-message
	Message EventType = "m.room.message"

	// https://matrix.org/docs/spec/client_server/latest#m-room-message-feedback
	MessageFeedback EventType = "m.room.message.feedback"

	// https://matrix.org/docs/spec/client_server/latest#m-room-name
	Name EventType = "m.room.name"

	// https://matrix.org/docs/spec/client_server/latest#m-room-topic
	Topic EventType = "m.room.topic"

	// https://matrix.org/docs/spec/client_server/latest#m-room-avatar
	Avatar EventType = "m.room.avatar"

	// https://matrix.org/docs/spec/client_server/latest#m-room-pinned-events
	PinnedEvents EventType = "m.room.pinned_events"
)

type Event interface {
	Content() json.RawMessage // Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Type() EventType          // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	ID() string
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

type signed struct {
	Mxid string `json:"mxid"` // Required. The invited matrix user ID. Must be equal to the user_id property of the event.
	// TODO:
	// Signatures Signatures `json:"signatures"` // Required. A single signature from the verifying server, in the format specified by the Signing Events section of the server-server API.
	Token string `json:"token"` // Required. The token property of the containing third_party_invite object.
}

type State struct {
	events []StateEvent `json:"events"` // List of events.
}

type Invite struct {
	DisplayName string `json:"display_name"` // Required. A name which can be displayed to represent the user instead of their third party identifier
	Signed      signed `json:"signed"`       // Required. A block of content which has been signed, which servers can use to verify the event. Clients should ignore this.
}

type Ephemeral struct {
	Events []Event `json:"events"` // List of events.
}

type StrippedState struct {
	// TODO: в документации EventContent, хотя вроде сервер выдаёт json.RawMessage
	Content  json.RawMessage `json:"content"`   // Required. The content for the event.
	StateKey string          `json:"state_key"` // Required. The state_key for the event.
	Type     string          `json:"type"`      // Required. The type for the event.
	Sender   string          `json:"sender"`    // Required. The sender for the event.
}

type UnsignedData struct {
	Age             int    `json:"age"`              // The time in milliseconds that has elapsed since the event was sent. This field is generated by the local homeserver, and may be incorrect if the local time on at least one of the two servers is out of sync, which can cause the age to either be negative or greater than it actually is.
	RedactedBecause Event  `json:"redacted_because"` // Optional. The event that redacted this event, if any.
	TransactionID   string `json:"transaction_id"`   // The client-supplied transaction ID, if the client being given the event is the same one which sent it.
}

type Presence struct {
	Events []Event `json:"events"` // List of events.
}

type ToDevice struct {
	Events []Event `json:"events` // List of send-to-device messages
}
