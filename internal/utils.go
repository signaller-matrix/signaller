package internal

import "github.com/signaller-matrix/signaller/internal/models/publicrooms"

func roomsToPublicRoomsChunks(rooms []Room) []publicrooms.PublicRoomsChunk {
	var chunks []publicrooms.PublicRoomsChunk

	for _, room := range rooms {
		chunk := publicrooms.PublicRoomsChunk{
			// TODO: Aliases:
			CanonicalAlias:   room.AliasName(),
			Name:             room.Name(),
			NumJoinedMembers: len(room.Users()),
			RoomID:           room.ID(),
			Topic:            room.Topic(),
			WorldReadable:    room.WorldReadable(),
			GuestCanJoin:     room.GuestCanJoin(),
			AvatarURL:        room.AvatarURL()}

		chunks = append(chunks, chunk)
	}

	return chunks
}
