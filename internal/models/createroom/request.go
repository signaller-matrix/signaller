package createroom

import (
	common "github.com/nxshock/signaller/internal/models/common"
)

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-createroom
type VisibilityType string

const (
	VisibilityTypePrivate VisibilityType = "private"
	VisibilityTypePublic  VisibilityType = "public"
)

type Preset string

const (
	PrivateChat        Preset = "private_chat"
	PublicChat         Preset = "public_chat"
	TrustedPrivateChat Preset = "trusted_private_chat"
)

// Invite3pid represents third party IDs to invite into the room
// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-createroom
type Invite3pid struct {
	IDServer string `json:"id_server"` // Required. The hostname+port of the identity server which should be used for third party identifier lookups.
	Medium   string `json:"medium"`    // Required. The kind of address being passed in the address field, for example email.
	Address  string `json:"address"`   // Required. The invitee's third party identifier.
}

// Request is room creation request
// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-createroom
type Request struct {
	Visibility    VisibilityType `json:"visibility,omitempty"`      // A public visibility indicates that the room will be shown in the published room list. A private visibility will hide the room from the published room list. Rooms default to private visibility if this key is not included. NB: This should not be confused with join_rules which also uses the word public. One of: ["public", "private"]
	RoomAliasName string         `json:"room_alias_name,omitempty"` // The desired room alias local part. If this is included, a room alias will be created and mapped to the newly created room. The alias will belong on the same homeserver which created the room. For example, if this was set to "foo" and sent to the homeserver "example.com" the complete room alias would be #foo:example.com.
	Name          string         `json:"name,omitempty"`            // If this is included, an m.room.name event will be sent into the room to indicate the name of the room. See Room Events for more information on m.room.name.
	Topic         string         `json:"topic,omitempty"`           // If this is included, an m.room.topic event will be sent into the room to indicate the topic for the room. See Room Events for more information on m.room.topic.
	Invite        []string       `json:"invite,omitempty"`          // A list of user IDs to invite to the room. This will tell the server to invite everyone in the list to the newly created room.
	Invite3pids   []Invite3pid   `json:"invite_3pid,omitempty"`     // A list of objects representing third party IDs to invite into the room.
	RoomVersion   string         `json:"room_version,omitempty"`    // The room version to set for the room. If not provided, the homeserver is to use its configured default. If provided, the homeserver will return a 400 error with the errcode M_UNSUPPORTED_ROOM_VERSION if it does not support the room version.
	// TODO: проверить тип
	// CreationContent CreationContentType `json:"creation_content,omitempty"`
	InitialState []common.StateEvent `json:"initial_state,omitempty"` // A list of state events to set in the new room. This allows the user to override the default state events set in the new room. The expected format of the state events are an object with type, state_key and content keys set. Takes precedence over events set by preset, but gets overriden by name and topic keys.
	Preset       Preset              `json:"preset,omitempty"`        // Convenience parameter for setting various default state events based on a preset. If unspecified, the server should use the visibility to determine which preset to use. A visbility of public equates to a preset of public_chat and private visibility equates to a preset of private_chat. One of: ["private_chat", "public_chat", "trusted_private_chat"]
	IsDirect     bool                `json:"is_direct,omitempty"`     // This flag makes the server set the is_direct flag on the m.room.member events sent to the users in invite and invite_3pid.
	// PowerLevelContentOverride `json:"power_level_content_override"`
}
