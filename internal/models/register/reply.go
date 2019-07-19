package register

// https://matrix.org/docs/spec/client_server/latest#post-matrix-client-r0-register
type RegisterResponse struct {
	UserID      string `json:"user_id"`                // Required. The fully-qualified Matrix user ID (MXID) that has been registered. Any user ID returned by this API must conform to the grammar given in the Matrix specification.
	AccessToken string `json:"access_token,omitempty"` // An access token for the account. This access token can then be used to authorize other requests. Required if the inhibit_login option is false.
	DeviceID    string `json:"device_id,omitempty"`    // ID of the registered device. Will be the same as the corresponding parameter in the request, if one was specified. Required if the inhibit_login option is false.
}
