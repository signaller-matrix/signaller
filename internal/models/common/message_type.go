package common

// https://matrix.org/docs/spec/client_server/r0.4.0.html#m-room-message-msgtypes
type MessageType string

const (
	MessageTypeText     MessageType = "m.text"     // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-text
	MessageTypeEmote    MessageType = "m.emote"    // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-emote
	MessageTypeNotice   MessageType = "m.notice"   // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-notice
	MessageTypeImage    MessageType = "m.image"    // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-image
	MessageTypeFile     MessageType = "m.file"     // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-file
	MessageTypeVideo    MessageType = "m.video"    // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-video
	MessageTypeAudio    MessageType = "m.audio"    // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-audio
	MessageTypeLocation MessageType = "m.location" // https://matrix.org/docs/spec/client_server/r0.4.0.html#m-location
)
