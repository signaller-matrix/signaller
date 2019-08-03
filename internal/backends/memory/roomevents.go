package memory

import (
	"encoding/json"
	"time"

	"github.com/signaller-matrix/signaller/internal"
	"github.com/signaller-matrix/signaller/internal/models/rooms"
)

type RoomEvent struct {
	Content        json.RawMessage
	Type           rooms.Type
	EventID        string
	Sender         internal.User
	OriginServerTS time.Time
	Room           internal.Room
}

func (roomEvent *RoomEvent) ToEvent() rooms.Event {
	event := rooms.Event{
		Content:        roomEvent.Content,
		Type:           roomEvent.Type,
		EventID:        roomEvent.EventID,
		Sender:         roomEvent.Sender.ID(),
		OriginServerTS: roomEvent.OriginServerTS.Unix(),
		RoomID:         roomEvent.Room.ID()}

	return event
}
