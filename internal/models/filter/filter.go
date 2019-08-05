package filter

type EventFormat string

const (
	EventFormatClient     = "client"
	EventFormatFederation = "federation"
)

type Request struct {
	EventFields []string    `json:"event_fields"` // List of event fields to include. If this list is absent then all fields are included. The entries may include '.' charaters to indicate sub-fields. So ['content.body'] will include the 'body' field of the 'content' object. A literal '.' character in a field name may be escaped using a '\'. A server may include more fields than were requested.
	EventFormat EventFormat `json:"event_format"` // The format to use for events. 'client' will return the events in a format suitable for clients. 'federation' will return the raw event as receieved over federation. The default is 'client'. One of: ["client", "federation"]
	Presence    EventFilter `json:"presence"`     // The presence updates to include.
	AccountData EventFilter `json:"account_data"` // The user account data that isn't associated with rooms to include.
	Room        RoomFilter  `json:"room"`         // Filters to be applied to room data.
}

type EventFilter struct {
	Limit      int      `json:"limit"`       // The maximum number of events to return.
	NotSenders []string `json:"not_senders"` // A list of sender IDs to exclude. If this list is absent then no senders are excluded. A matching sender will be excluded even if it is listed in the 'senders' filter.
	NotTypes   []string `json:"not_types"`   // A list of event types to exclude. If this list is absent then no event types are excluded. A matching type will be excluded even if it is listed in the 'types' filter. A '*' can be used as a wildcard to match any sequence of characters.
	Senders    []string `json:"senders"`     // A list of senders IDs to include. If this list is absent then all senders are included.
	Types      []string `json:"types"`       // A list of event types to include. If this list is absent then all event types are included. A '*' can be used as a wildcard to match any sequence of characters.
}

type RoomFilter struct {
	NotRooms     []string        `json:"not_rooms"`     // A list of room IDs to exclude. If this list is absent then no rooms are excluded. A matching room will be excluded even if it is listed in the 'rooms' filter. This filter is applied before the filters in ephemeral, state, timeline or account_data
	Rooms        []string        `json:"rooms"`         // A list of room IDs to include. If this list is absent then all rooms are included. This filter is applied before the filters in ephemeral, state, timeline or account_data
	Ephemeral    RoomEventFilter `json:"ephemeral"`     // The events that aren't recorded in the room history, e.g. typing and receipts, to include for rooms.
	IncludeLeave bool            `json:"include_leave"` // Include rooms that the user has left in the sync, default false
	State        StateFilter     `json:"state"`         // The state events to include for rooms.
	Timeline     RoomEventFilter `json:"timeline"`      // The message and state update events to include for rooms.
	AccountData  RoomEventFilter `json:"account_data"`  // The per user account data to include for rooms.
}

type StateFilter struct {
	Limit                   int      `json:"limit"`                   // The maximum number of events to return.
	NotSenders              []string `json:"notSenders"`              // A list of sender IDs to exclude. If this list is absent then no senders are excluded. A matching sender will be excluded even if it is listed in the 'senders' filter.
	NotTypes                []string `json:"notTypes"`                // A list of event types to exclude. If this list is absent then no event types are excluded. A matching type will be excluded even if it is listed in the 'types' filter. A '*' can be used as a wildcard to match any sequence of characters.
	Senders                 []string `json:"senders"`                 // A list of senders IDs to include. If this list is absent then all senders are included.
	Types                   []string `json:"types"`                   // A list of event types to include. If this list is absent then all event types are included. A '*' can be used as a wildcard to match any sequence of characters.
	LazyLoadMembers         bool     `json:"lazyLoadMembers"`         // If true, enables lazy-loading of membership events. See Lazy-loading room members for more information. Defaults to false.
	IncludeRedundantMembers bool     `json:"includeRedundantMembers"` // If true, sends all membership events for all events, even if they have already been sent to the client. Does not apply unless lazyLoadMembers is true. See Lazy- loading room members for more information. Defaults to false.
	NotRooms                []string `json:"notRooms"`                // A list of room IDs to exclude. If this list is absent then no rooms are excluded. A matching room will be excluded even if it is listed in the 'rooms' filter.
	Rooms                   []string `json:"rooms"`                   // A list of room IDs to include. If this list is absent then all rooms are included.
	Contains_url            bool     `json:"contains_url"`            // If true, includes only events with a url key in their content. If false, excludes those events. If omitted, url key is not considered for filtering.
}

type RoomEventFilter struct {
	Limit                   int      `json:"limit"`                   // The maximum number of events to return.
	NotSenders              []string `json:"notSenders"`              // A list of sender IDs to exclude. If this list is absent then no senders are excluded. A matching sender will be excluded even if it is listed in the 'senders' filter.
	NotTypes                []string `json:"notTypes"`                // A list of event types to exclude. If this list is absent then no event types are excluded. A matching type will be excluded even if it is listed in the 'types' filter. A '*' can be used as a wildcard to match any sequence of characters.
	Senders                 []string `json:"senders"`                 // A list of senders IDs to include. If this list is absent then all senders are included.
	Types                   []string `json:"types"`                   // A list of event types to include. If this list is absent then all event types are included. A '*' can be used as a wildcard to match any sequence of characters.
	LazyLoadMembers         bool     `json:"lazyLoadMembers"`         // If true, enables lazy-loading of membership events. See Lazy-loading room members for more information. Defaults to false.
	IncludeRedundantMembers bool     `json:"includeRedundantMembers"` // If true, sends all membership events for all events, even if they have already been sent to the client. Does not apply unless lazyLoadMembers is true. See Lazy- loading room members for more information. Defaults to false.
	NotRooms                []string `json:"notRooms"`                // A list of room IDs to exclude. If this list is absent then no rooms are excluded. A matching room will be excluded even if it is listed in the 'rooms' filter.
	Rooms                   []string `json:"rooms"`                   // A list of room IDs to include. If this list is absent then all rooms are included.
	Contains_url            bool     `json:"contains_url"`            // If true, includes only events with a url key in their content. If false, excludes those events. If omitted, url key is not considered for filtering.
}

type Response struct {
	FilterID string `json:"filter_id"` // Required. The ID of the filter that was created. Cannot start with a { as this character is used to determine if the filter provided is inline JSON or a previously declared filter by homeservers on some APIs.
}
