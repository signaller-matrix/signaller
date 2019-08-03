package login

import (
	"github.com/signaller-matrix/signaller/internal/models/common"
)

// PostRequest represents POST login request
// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-login
type PostRequest struct {
	Type                     common.AuthenticationType `json:"type"`                                  // Required. The login type being used. One of: ["m.login.password", "m.login.token"]
	Identifier               common.UserIdentifier     `json:"identifier"`                            // Identification information for the user.
	Password                 string                    `json:"password,omitempty"`                    // Required when type is m.login.password. The user's password.
	Token                    string                    `json:"token,omitempty"`                       // Required when type is m.login.token. Part of Token-based login.
	DeviceID                 string                    `json:"device_id,omitempty"`                   // ID of the client device. If this does not correspond to a known client device, a new device will be created. The server will auto-generate a device_id if this is not specified.
	InitialDeviceDisplayName string                    `json:"initial_device_display_name,omitempty"` // A display name to assign to the newly-created device. Ignored if device_id corresponds to a known device.
}
