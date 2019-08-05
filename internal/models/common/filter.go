package common

type Filter struct {
	Room struct {
		State struct {
			Types    []string `json:"types"`
			NotRooms []string `json:"not_rooms"`
		} `json:"state"`
		Timeline struct {
			Limit      int      `json:"limit"`
			Types      []string `json:"types"`
			NotRooms   []string `json:"not_rooms"`
			NotSenders []string `json:"not_senders"`
		} `json:"timeline"`
		Ephemeral struct {
			Types      []string `json:"types"`
			NotRooms   []string `json:"not_rooms"`
			NotSenders []string `json:"not_senders"`
		} `json:"ephemeral"`
	} `json:"room"`
	Presence struct {
		Types      []string `json:"types"`
		NotSenders []string `json:"not_senders"`
	} `json:"presence"`
	EventFormat string   `json:"event_format"`
	EventFields []string `json:"event_fields"`
}
