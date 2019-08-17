package events

import (
	"encoding/json"
)

type RoomEvent struct {
	// TODO: object
	ContentData    json.RawMessage `json:"content"`          // Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	EType          EventType       `json:"type"`             // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
	EventID        string          `json:"event_id"`         // Required. The globally unique event identifier.
	Sender         string          `json:"sender"`           // Required. Contains the fully-qualified ID of the user who sent this event.
	OriginServerTs int64           `json:"origin_server_ts"` // Required. Timestamp in milliseconds on originating homeserver when this event was sent.
	Unsigned       UnsignedData    `json:"unsigned"`         // Contains optional extra information about the event.
	RoomID         string          `json:"room_id"`          // Required. The ID of the room associated with this event. Will not be present on events that arrive through /sync, despite being required everywhere else.
}

func (roomEvent RoomEvent) Content() json.RawMessage {
	return roomEvent.ContentData
}

func (roomEvent RoomEvent) ID() string {
	return roomEvent.EventID
}

func (roomEvent RoomEvent) Type() EventType {
	return roomEvent.EType
}
