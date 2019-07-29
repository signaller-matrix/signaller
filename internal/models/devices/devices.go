package devices

type Response struct {
	Devices []Device `json:"devices"` // A list of all registered devices for this user.
}

type Device struct {
	DeviceID    string `json:"device_id"`    // Required. Identifier of this device.
	DisplayName string `json:"display_name"` // Display name set by the user for this device. Absent if no name has been set.
	LastSeenIP  string `json:"last_seen_ip"` // The IP address where this device was last seen. (May be a few minutes out of date, for efficiency reasons).
	LastSeenTS  int    `json:"last_seen_ts"` // The timestamp (in milliseconds since the unix epoch) when this devices was last seen. (May be a few minutes out of date, for efficiency reasons).
}
