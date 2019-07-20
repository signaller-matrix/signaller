package events

import (
	"encoding/json"
)

// Type is type of event
type Type string

const (
	// https://matrix.org/docs/spec/client_server/latest#m-room-aliases
	RoomAliases Type = "m.room.aliases"

	// https://matrix.org/docs/spec/client_server/latest#m-room-canonical-alias
	RoomCanonicalAlias Type = "m.room.canonical_alias"

	// https://matrix.org/docs/spec/client_server/latest#m-room-create
	RoomCreate Type = "m.room.create"

	// https://matrix.org/docs/spec/client_server/latest#m-room-join-rules
	RoomJoinRules Type = "m.room.join_rules"

	// https://matrix.org/docs/spec/client_server/latest#m-room-member
	RoomMember Type = "m.room.member"

	// https://matrix.org/docs/spec/client_server/latest#m-room-power-levels
	RoomPowerLevels Type = "m.room.power_levels"

	// https://matrix.org/docs/spec/client_server/latest#m-room-redaction
	RoomRedaction Type = "m.room.redaction"
)

// Event is the basic set of fields all events must have
// https://matrix.org/docs/spec/client_server/latest#event-fields
type Event struct {
	Content json.RawMessage `json:"content"` //	Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Type    Type            `json:"type"`    // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
}
