package sync

import (
	"github.com/signaller-matrix/signaller/internal/models/events"
)

type SyncReply struct {
	NextBatch              string             `json:"next_batch"`                 // Required. The batch token to supply in the since param of the next /sync request.
	Rooms                  RoomsSyncReply     `json:"rooms"`                      // Updates to rooms.
	Presence               events.Presence    `json:"presence"`                   // The updates to the presence status of other users.
	AccountData            events.AccountData `json:"account_data"`               // The global private data created by this user.
	ToDevice               events.ToDevice    `json:"to_device"`                  // Information on the send-to-device messages for the client device, as defined in Send-to-Device messaging.
	DeviceLists            events.DeviceLists `json:"device_lists"`               // Information on end-to-end device updates, as specified in End-to-end encryption.
	DeviceOneTimeKeysCount map[string]int     `json:"device_one_time_keys_count"` // Information on end-to-end encryption keys, as specified in End-to-end encryption.
}

type RoomsSyncReply struct {
	Join   map[string]events.JoinedRoom  `json:"join"`   // The rooms that the user has joined.
	Invite map[string]events.InvitedRoom `json:"invite"` // The rooms that the user has been invited to.
	Leave  map[string]events.LeftRoom    `json:"leave"`  // The rooms that the user has left or been banned from.
}
