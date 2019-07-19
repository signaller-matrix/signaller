package models

import (
	common "github.com/nxshock/signaller/internal/models/common"
)

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-versions
type VersionsReply struct {
	Versions         []string        `json:"versions"` // The supported versions.
	UnstableFeatures map[string]bool `json:"unstable_features,omitempty"`
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-account-whoami
type WhoAmIReply struct {
	UserID string `json:"user_id"` // Required. The user id that owns the access token.
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#get-matrix-client-r0-rooms-roomid-members
type MembersReply struct {
	Chunk []common.MemberEvent `json:"chunk"`
}
