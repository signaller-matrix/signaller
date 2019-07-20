package versions

// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-versions
type Reply struct {
	Versions         []string        `json:"versions"`                    // The supported versions.
	UnstableFeatures map[string]bool `json:"unstable_features,omitempty"` // Experimental features the server supports. Features not listed here, or the lack of this property all together, indicate that a feature is not supported.
}
