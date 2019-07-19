package common

type DeviceLists struct {
	Changed []string `json:"changed"` // List of users who have updated their device identity keys, or who now share an encrypted room with the client since the previous sync response.
	Left    []string `json:"left"`    // List of users with whom we do not share any encrypted rooms anymore since the previous sync response.
}
