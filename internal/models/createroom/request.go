package createroom

import (
	common "github.com/nxshock/signaller/internal/models/common"
)

type VisibilityType string

const (
	VisibilityTypePrivate = "private"
	VisibilityTypePublic  = "public"
)

// Invite3pid represents third party IDs to invite into the room
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-createroom
type Invite3pid struct {
	IDServer string `json:"id_server"` // Required. The hostname+port of the identity server which should be used for third party identifier lookups.
	Medium   string `json:"medium"`    // Required. The kind of address being passed in the address field, for example email.
	Address  string `json:"address"`   // Required. The invitee's third party identifier.
}

// CreateRoomRequest represents room creation request
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-createroom
type CreateRoomRequest struct {
	Visibility    VisibilityType `json:"visibility,omitempty"`
	RoomAliasName string         `json:"room_alias_name,omitempty"`
	Name          string         `json:"name,omitempty"`
	Topic         string         `json:"topic,omitempty"`
	Invite        []string       `json:"invite,omitempty"`
	Invite3pids   []Invite3pid   `json:"invite_3pid,omitempty"`
	RoomVersion   string         `json:"room_version,omitempty"`
	// TODO: проверить тип
	// CreationContent CreationContentType `json:"creation_content,omitempty"`
	InitialState []common.StateEvent `json:"initial_state,omitempty"`
	Preset       string              `json:"preset,omitempty"` // TODO: проверить тип
	IsDirect     bool                `json:"is_direct,omitempty"`
	// PowerLevelContentOverride `json:"power_level_content_override"`
}
