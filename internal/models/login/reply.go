package login

import (
	common "github.com/signaller-matrix/signaller/internal/models/common"
)

// PostReply is returned reply from POST login method
// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login
type PostReply struct {
	UserID      string               `json:"user_id"`              // The fully-qualified Matrix ID that has been registered.
	AccessToken string               `json:"access_token"`         // An access token for the account. This access token can then be used to authorize other requests.
	DeviceID    string               `json:"device_id"`            // ID of the logged-in device. Will be the same as the corresponding parameter in the request, if one was specified.
	WellKnown   DiscoveryInformation `json:"well_known,omitempty"` // Optional client configuration provided by the server. If present, clients SHOULD use the provided object to reconfigure themselves, optionally validating the URLs within. This object takes the same form as the one returned from .well-known autodiscovery.
}

// DiscoveryInformation is client configuration provided by the server
// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login
type DiscoveryInformation struct {
	HomeServer     HomeserverInformation     `json:"m.homeserver"`                // Required. Used by clients to discover homeserver information.
	IdentityServer IdentityServerInformation `json:"m.identity_server,omitempty"` // Used by clients to discover ide
}

// HomeserverInformation is used by clients to discover homeserver information
// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login
type HomeserverInformation struct {
	BaseURL string `json:"base_url"` // Required. The base URL for the homeserver for client-server connections.
}

// IdentityServerInformation is used by clients to discover identity server information.
// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login
type IdentityServerInformation struct {
	BaseURL string `json:"base_url"` // Required. The base URL for the homeserver for client-server connections.
}

// Flow is the homeserver's supported login types
// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login
type Flow struct {
	Type common.AuthenticationType `json:"type"`
}

// GetReply is returned reply from GET login method
// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login
type GetReply struct {
	Flows []Flow `json:"flows"` // The homeserver's supported login types
}
