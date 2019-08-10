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

func (roomEvent *RoomEvent) Content() json.RawMessage {
	return roomEvent.ContentData
}

func (roomEvent *RoomEvent) ID() string {
	return roomEvent.EventID
}

func (roomEvent *RoomEvent) Type() EventType {
	return roomEvent.EType
}

type JoinedRoom struct {
	State               State                    `json:"state"`                // Updates to the state, between the time indicated by the since parameter, and the start of the timeline (or all state up to the start of the timeline, if since is not given, or full_state is true).
	Timeline            Timeline                 `json:"timeline"`             // The timeline of messages and state changes in the room.
	Ephemeral           Ephemeral                `json:"ephemeral"`            // The ephemeral events in the room that aren't recorded in the timeline or state of the room. e.g. typing.
	AccountData         AccountData              `json:"account_data"`         // The private data that this user has attached to this room.
	UnreadNotifications UnreadNotificationCounts `json:"unread_notifications"` // Counts of unread notifications for this room
}

type AccountData struct {
	Events []Event `json:"events"` // List of events.
}

type UnreadNotificationCounts struct {
	HighlightCount    int `json:"highlight_count"`    // The number of unread notifications for this room with the highlight flag set
	NotificationCount int `json:"notification_count"` // The total number of unread notifications for this room
}

type LeftRoom struct {
	State       State       `json:"state"`        // The state updates for the room up to the start of the timeline.
	Timeline    Timeline    `json:"timeline"`     // The timeline of messages and state changes in the room up to the point when the user left.
	AccountData AccountData `json:"account_data"` // The private data that this user has attached to this room.
}

type InvitedRoom struct {
	InviteState InviteState `json:"invite_state"` // The state of a room that the user has been invited to. These state events may only have the sender, type, state_key and content keys present. These events do not replace any state that the client already has for the room, for example if the client has archived the room. Instead the client should keep two separate copies of the state: the one from the invite_state and one from the archived state. If the client joins the room then the current state will be given as a delta against the archived state not the invite_state.
}

type InviteState struct {
	Events []StrippedState `json:"events"` // The StrippedState events that form the invite state.
}

type Timeline struct {
	Events    []RoomEvent `json:"events"`     // List of events.
	Limited   bool        `json:"limited"`    // True if the number of events returned was limited by the limit on the filter.
	PrevBatch string      `json:"prev_batch"` // A token that can be supplied to the from parameter of the rooms/{roomId}/messages endpoint.
}
