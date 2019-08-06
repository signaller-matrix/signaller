package roomalias

// Merged PUT and GET requests
type Request struct {
	RoomID    string `json:"room_id"`   // The room ID to set.
	RoomAlias string `json:"roomAlias"` // The room alias.
}

type ResponseGet struct {
	RoomID  string   `json:"room_id"` // The room ID for this room alias.
	Servers []string `json:"servers"` // A list of servers that are aware of this room alias.
}
