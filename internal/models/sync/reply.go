package sync

import (
	"github.com/signaller-matrix/signaller/internal/models/events"
)

// SyncReply is model of sync respond
type SyncReply struct {
	NextBatch              string             `json:"next_batch"`                 // Required. The batch token to supply in the since param of the next /sync request.
	Rooms                  RoomsSyncReply     `json:"rooms"`                      // Updates to rooms.
	Presence               events.Presence    `json:"presence"`                   // The updates to the presence status of other users.
	AccountData            AccountData        `json:"account_data"`               // The global private data created by this user.
	ToDevice               events.ToDevice    `json:"to_device"`                  // Information on the send-to-device messages for the client device, as defined in Send-to-Device messaging.
	DeviceLists            events.DeviceLists `json:"device_lists"`               // Information on end-to-end device updates, as specified in End-to-end encryption.
	DeviceOneTimeKeysCount map[string]int     `json:"device_one_time_keys_count"` // Information on end-to-end encryption keys, as specified in End-to-end encryption.
}

type RoomsSyncReply struct {
	Join   map[string]JoinedRoom  `json:"join"`   // The rooms that the user has joined.
	Invite map[string]InvitedRoom `json:"invite"` // The rooms that the user has been invited to.
	Leave  map[string]LeftRoom    `json:"leave"`  // The rooms that the user has left or been banned from.
}

type RoomSummary struct {
	Heroes             []string `json:"m.heroes"`
	JoinedMemberCount  int      `json:"m.joined_member_count"`
	InvitedMemberCount int      `json:"m.invited_member_count"`
}

type JoinedRoom struct {
	RoomSummary         RoomSummary              `json:"summary"`
	State               events.State             `json:"state"`                // Updates to the state, between the time indicated by the since parameter, and the start of the timeline (or all state up to the start of the timeline, if since is not given, or full_state is true).
	Timeline            Timeline                 `json:"timeline"`             // The timeline of messages and state changes in the room.
	Ephemeral           events.Ephemeral         `json:"ephemeral"`            // The ephemeral events in the room that aren't recorded in the timeline or state of the room. e.g. typing.
	AccountData         AccountData              `json:"account_data"`         // The private data that this user has attached to this room.
	UnreadNotifications UnreadNotificationCounts `json:"unread_notifications"` // Counts of unread notifications for this room
}

type AccountData struct {
	Events []events.Event `json:"events"` // List of events.
}

type UnreadNotificationCounts struct {
	HighlightCount    int `json:"highlight_count"`    // The number of unread notifications for this room with the highlight flag set
	NotificationCount int `json:"notification_count"` // The total number of unread notifications for this room
}

type LeftRoom struct {
	State       events.State `json:"state"`        // The state updates for the room up to the start of the timeline.
	Timeline    Timeline     `json:"timeline"`     // The timeline of messages and state changes in the room up to the point when the user left.
	AccountData AccountData  `json:"account_data"` // The private data that this user has attached to this room.
}

type InvitedRoom struct {
	InviteState InviteState `json:"invite_state"` // The state of a room that the user has been invited to. These state events may only have the sender, type, state_key and content keys present. These events do not replace any state that the client already has for the room, for example if the client has archived the room. Instead the client should keep two separate copies of the state: the one from the invite_state and one from the archived state. If the client joins the room then the current state will be given as a delta against the archived state not the invite_state.
}

type InviteState struct {
	Events []events.StrippedState `json:"events"` // The StrippedState events that form the invite state.
}

type Timeline struct {
	Events    []events.RoomEvent `json:"events"`     // List of events.
	Limited   bool               `json:"limited"`    // True if the number of events returned was limited by the limit on the filter.
	PrevBatch string             `json:"prev_batch"` // A token that can be supplied to the from parameter of the rooms/{roomId}/messages endpoint.
}

// BuildEmptySyncReply is function which builds empty SyncReply model
func BuildEmptySyncReply() *SyncReply {
	return &SyncReply{
		NextBatch: "",
		Rooms: RoomsSyncReply{
			Join:   make(map[string]JoinedRoom),
			Invite: make(map[string]InvitedRoom),
			Leave:  make(map[string]LeftRoom),
		},
		Presence: events.Presence{
			Events: nil,
		},
		AccountData: AccountData{
			Events: nil,
		},
		ToDevice: events.ToDevice{
			Events: nil,
		},
		DeviceLists: events.DeviceLists{
			Changed: nil,
			Left:    nil,
		},
		DeviceOneTimeKeysCount: make(map[string]int),
	}
}
