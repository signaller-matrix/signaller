package internal

import (
	"strings"

	"github.com/signaller-matrix/signaller/internal/models/publicrooms"
)

func GetCanonicalAlias(hostName string, alias string) string {
	return "#" + alias + ":" + hostName
}

func StripAlias(hostName string, canonicalAlias string) string {
	canonicalAlias = strings.TrimPrefix(canonicalAlias, "#")
	canonicalAlias = strings.TrimSuffix(canonicalAlias, ":"+hostName)

	return canonicalAlias
}

func roomsToPublicRoomsChunks(rooms []Room) []publicrooms.PublicRoomsChunk {
	var chunks []publicrooms.PublicRoomsChunk

	for _, room := range rooms {
		chunk := publicrooms.PublicRoomsChunk{
			Aliases:          room.Aliases(),
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
