package capabilities

// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-capabilities
type Response struct {
	Capabilities Capabilities `json:"capabilities"` // Required. The custom capabilities the server supports, using the Java package naming convention.
}

type Capabilities struct {
	ChangePassword ChangePasswordCapability `json:"m.change_password"` // Capability to indicate if the user can change their password.
	RoomVersions   RoomVersionsCapability   `json:"m.room_versions"`   // The room versions the server supports.
}

type ChangePasswordCapability struct {
	Enabled bool `json:"enabled"` // Required. True if the user can change their password, false otherwise.
}

type RoomVersionsCapability struct {
	Default   string               `json:"default"`   // Required. The default room version the server is using for new rooms.
	Available map[string]Stability `json:"available"` // Required. A detailed description of the room versions the server supports.
}

type Stability string

const (
	Stable   Stability = "stable"
	Unstable Stability = "unstable"
)
