package password

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-account-password
type Request struct {
	NewPassword string             `json:"new_password"`   // Required. The new password for the account.
	Auth        AuthenticationData `json:"auth,omitempty"` // Additional authentication information for the user-interactive authentication API.
}

type AuthenticationData struct {
	Type    string `json:"type"`    // Required. The login type that the client is attempting to complete.
	Session string `json:"session"` //The value of the session key given by the homeserver.
}
