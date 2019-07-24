package joinedrooms

// https://matrix.org/docs/spec/client_server/latest#get-matrix-client-r0-joined-rooms
type Response struct {
	JoinedRooms []string `json:"joined_rooms"` // Required. The ID of each room in which the user has joined membership.
}
