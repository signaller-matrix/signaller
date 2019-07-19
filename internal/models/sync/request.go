package sync

type SetPresence string

const (
	SetPresenceOffline     SetPresence = "offline"
	SetPresenceOnline      SetPresence = "online"
	SetPresenceUnavailable SetPresence = "unavailable"
)

type SyncRequest struct {
	Filter      string      `url:"filter,omitempty"`       // The ID of a filter created using the filter API or a filter JSON object encoded as a string. The server will detect whether it is an ID or a JSON object by whether the first character is a "{" open brace. Passing the JSON inline is best suited to one off requests. Creating a filter using the filter API is recommended for clients that reuse the same filter multiple times, for example in long poll requests.
	Since       string      `url:"since,omitempty"`        // A point in time to continue a sync from.
	FullState   bool        `url:"full_state,omitempty"`   // Controls whether to include the full state for all rooms the user is a member of.
	SetPresence SetPresence `url:"set_presence,omitempty"` // Controls whether the client is automatically marked as online by polling this API. If this parameter is omitted then the client is automatically marked as online when it uses this API. Otherwise if the parameter is set to "offline" then the client is not marked as being online when it uses this API. When set to "unavailable", the client is marked as being idle. One of: ["offline", "online", "unavailable"]
	Timeout     int         `url:"timeout,omitempty"`      // The maximum time to wait, in milliseconds, before returning this request. If no events (or other data) become available before this time elapses, the server will return a response with empty fields.
}
