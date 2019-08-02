package publicrooms

type Request struct {
	Limit  int    `json:"limit"`  // Limit the number of results returned.
	Since  string `json:"since"`  // A pagination token from a previous request, allowing clients to get the next (or previous) batch of rooms. The direction of pagination is specified solely by which token is supplied, rather than via an explicit flag.
	Server string `json:"server"` // The server to fetch the public room lists from. Defaults to the local server.
}

type Response struct {
	Chunk                  []PublicRoomsChunk `json:"chunk"`                               // A paginated chunk of public rooms.
	NextBatch              string             `json:"next_batch,omitempty"`                // A pagination token for the response. The absence of this token means there are no more results to fetch and the client should stop paginating.
	PrevBatch              string             `json:"prev_batch,omitempty"`                // A pagination token that allows fetching previous results. The absence of this token means there are no results before this batch, i.e. this is the first batch.
	TotalRoomCountEstimate int                `json:"total_room_count_estimate,omitempty"` // An estimate on the total number of public rooms, if the server has an estimate.
}

type PublicRoomsChunk struct {
	Aliases          []string `json:"aliases,omitempty"`         // Aliases of the room. May be empty.
	CanonicalAlias   string   `json:"canonical_alias,omitempty"` // The canonical alias of the room, if any.
	Name             string   `json:"name,omitempty"`            // The name of the room, if any.
	NumJoinedMembers int      `json:"num_joined_members"`        // Required. The number of members joined to the room.
	RoomID           string   `json:"room_id"`                   // Required. The ID of the room.
	Topic            string   `json:"topic,omitempty"`           // The topic of the room, if any.
	WorldReadable    bool     `json:"world_readable"`            // Required. Whether the room may be viewed by guest users without joining.
	GuestCanJoin     bool     `json:"guest_can_join"`            // Required. Whether guest users may join the room and participate in it. If they can, they will be subject to ordinary power level rules like any other user.
	AvatarURL        string   `json:"avatar_url,omitempty"`      // The URL for the room's avatar, if one is set.
}
