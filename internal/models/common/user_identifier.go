package common

// UserIdentifier represents user identifier object
// https://matrix.org/docs/spec/client_server/r0.4.0.html#post-matrix-client-r0-login
type UserIdentifier struct {
	Type    IdentifierType `json:"type"`              // Required. The type of identification. See Identifier types for supported values and additional property descriptions.
	User    string         `json:"user,omitempty"`    // The fully qualified user ID or just local part of the user ID, to log in.
	Medium  string         `json:"medium,omitempty"`  // When logging in using a third party identifier, the medium of the identifier. Must be 'email'.
	Address string         `json:"address,omitempty"` // Third party identifier for the user.
	Country string         `json:"country,omitempty"`
	Phone   string         `json:"phone,omitempty"`
}

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id207
type IdentifierType string

const (
	IdentifierTypeUser       IdentifierType = "m.id.user"       // https://matrix.org/docs/spec/client_server/r0.4.0.html#id208
	IdentifierTypeThirdparty IdentifierType = "m.id.thirdparty" // https://matrix.org/docs/spec/client_server/r0.4.0.html#id209
	IdentifierTypePhone      IdentifierType = "m.id.phone"      // https://matrix.org/docs/spec/client_server/r0.4.0.html#id210
)
