package register

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-register
type RegisterRequest struct {
	Auth                     AuthenticationData `json:"auth"`                        // Additional authentication information for the user-interactive authentication API. Note that this information is not used to define how the registered user should be authenticated, but is instead used to authenticate the register call itself.
	BindEmail                bool               `json:"bind_email"`                  // If true, the server binds the email used for authentication to the Matrix ID with the identity server.
	BindMsisdn               bool               `json:"bind_msisdn"`                 // If true, the server binds the phone number used for authentication to the Matrix ID with the identity server.
	Username                 string             `json:"username"`                    // The basis for the localpart of the desired Matrix ID. If omitted, the homeserver MUST generate a Matrix ID local part.
	Password                 string             `json:"password"`                    // The desired password for the account.
	DeviceID                 string             `json:"device_id"`                   // ID of the client device. If this does not correspond to a known client device, a new device will be created. The server will auto-generate a device_id if this is not specified.
	InitialDeviceDisplayName string             `json:"initial_device_display_name"` // A display name to assign to the newly-created device. Ignored if device_id corresponds to a known device.
	InhibitLogin             bool               `json:"inhibit_login"`               // If true, an access_token and device_id should not be returned from this call, therefore preventing an automatic login. Defaults to false.

}

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-register
type AuthenticationData struct {
	Type    string `json:"type"`              // Required. The login type that the client is attempting to complete.
	Session string `json:"session,omitempty"` // The value of the session key given by the homeserver.
}
