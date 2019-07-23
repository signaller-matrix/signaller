package whoami

// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-account-whoami
type Response struct {
	UserID string `json:"user_id"` // Required. The user id that owns the access token.
}
