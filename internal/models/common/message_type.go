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

// https://matrix.org/docs/spec/client_server/latest#m-text
type MessageTextContent struct {
	Body          string      `json:"body"`                     // Required. The body of the message.
	Msgtype       MessageType `json:"msgtype"`                  // Required. Must be 'm.text'.
	Format        string      `json:"format,omitempty"`         // The format used in the formatted_body. Currently only org.matrix.custom.html is supported.
	FormattedBody string      `json:"formatted_body,omitempty"` // The formatted version of the body. This is required if format is specified.
}

// https://matrix.org/docs/spec/client_server/latest#m-emote
type MessageEmoteContent struct {
	Body          string      `json:"body"`                     // Required. The emote action to perform.
	Msgtype       MessageType `json:"msgtype"`                  // Required. Must be 'm.emote'.
	Format        string      `json:"format,omitempty"`         // The format used in the formatted_body. Currently only org.matrix.custom.html is supported.
	FormattedBody string      `json:"formatted_body,omitempty"` // The formatted version of the body. This is required if format is specified.
}

// https://matrix.org/docs/spec/client_server/latest#m-notice
type MessageNoticeContent struct {
	Body    string      `json:"body"`    // Required. The notice text to send.
	Msgtype MessageType `json:"msgtype"` // Required. Must be 'm.notice'.
}

// https://matrix.org/docs/spec/client_server/latest#m-image
type MessageImageContent struct {
	Body    string        `json:"body"`           // Required. A textual representation of the image. This could be the alt text of the image, the filename of the image, or some kind of content description for accessibility e.g. 'image attachment'.
	Info    ImageInfo     `json:"info,omitempty"` // Metadata about the image referred to in url.
	Msgtype MessageType   `json:"msgtype"`        // Required. Must be 'm.image'.
	URL     string        `json:"url"`            // Required. Required if the file is unencrypted. The URL (typically MXC URI) to the image.
	File    EncryptedFile `json:"file"`           // Required if the file is encrypted. Information on the encrypted file, as specified in End-to-end encryption.
}

type ImageInfo struct {
	H             int           `json:"h,omitempty"`              // The intended display height of the image in pixels. This may differ from the intrinsic dimensions of the image file.
	W             int           `json:"w,omitempty"`              // The intended display width of the image in pixels. This may differ from the intrinsic dimensions of the image file.
	Mimetype      string        `json:"mimetype,omitempty"`       // The mimetype of the image, e.g. image/jpeg.
	Size          int           `json:"size,omitempty"`           // Size of the image in bytes.
	ThumbnailURL  string        `json:"thumbnail_url,omitempty"`  // The URL (typically MXC URI) to a thumbnail of the image. Only present if the thumbnail is unencrypted.
	ThumbnailFile EncryptedFile `json:"thumbnail_file,omitempty"` // Information on the encrypted thumbnail file, as specified in End-to-end encryption. Only present if the thumbnail is encrypted.
	ThumbnailInfo ThumbnailInfo `json:"thumbnail_info,omitempty"` // Metadata about the image referred to in thumbnail_url.
}

type ThumbnailInfo struct {
	H        int    `json:"h,omitempty"`        // The intended display height of the image in pixels. This may differ from the intrinsic dimensions of the image file.
	W        int    `json:"w,omitempty"`        // The intended display width of the image in pixels. This may differ from the intrinsic dimensions of the image file.
	MimeType string `json:"mimetype,omitempty"` // The mimetype of the image, e.g. image/jpeg.
	Size     int    `json:"size,omitempty"`     // Size of the image in bytes.
}

// https://matrix.org/docs/spec/client_server/latest#m-file
type MessageFileContent struct {
	Body     string        `json:"body"`               // Required. A human-readable description of the file. This is recommended to be the filename of the original upload.
	filename string        `json:"filename,omitempty"` // The original filename of the uploaded file.
	info     FileInfo      `json:"info,omitempty"`     // Information about the file referred to in url.
	msgtype  string        `json:"msgtype"`            // Required. Must be 'm.file'.
	URL      string        `json:"url"`                // 	Required. Required if the file is unencrypted. The URL (typically MXC URI) to the file.
	file     EncryptedFile `json:"file"`               // 	Required if the file is encrypted. Information on the encrypted file, as specified in End-to-end encryption.
}

type FileInfo struct {
	MimeType      string        `json:"mimetype,omitempty"`       // The mimetype of the file e.g. application/msword.
	Size          int           `json:"size,omitempty"`           // The size of the file in bytes.
	ThumbnailURL  string        `json:"thumbnail_url,omitempty"`  // The URL to the thumbnail of the file. Only present if the thumbnail is unencrypted.
	ThumbnailFile EncryptedFile `json:"thumbnail_file,omitempty"` // Information on the encrypted thumbnail file, as specified in End-to-end encryption. Only present if the thumbnail is encrypted.
	ThumbnailInfo ThumbnailInfo `json:"thumbnail_info,omitempty"` // Metadata about the image referred to in thumbnail_url.
}

// https://matrix.org/docs/spec/client_server/latest#m-audio
type MessageAudioContent struct {
	Body    string        `json:"body"`           // Required. A description of the audio e.g. 'Bee Gees - Stayin' Alive', or some kind of content description for accessibility e.g. 'audio attachment'.
	Info    AudioInfo     `json:"info,omitempty"` // Metadata for the audio clip referred to in url.
	Msgtype string        `json:"msgtype"`        // Required. Must be 'm.audio'.
	URL     string        `json:"url"`            // Required. Required if the file is not encrypted. The URL (typically MXC URI) to the audio clip.
	File    EncryptedFile `json:"file"`           // Required if the file is encrypted. Information on the encrypted file, as specified in End-to-end encryption.
}

type AudioInfo struct {
	Duration int    `json:"duration,omitempty"` // The duration of the audio in milliseconds.
	Mimetype string `json:"mimetype,omitempty"` // The mimetype of the audio e.g. audio/aac.
	Size     int    `json:"size,omitempty"`     // The size of the audio clip in bytes.
}

// https://matrix.org/docs/spec/client_server/latest#m-location
type MessageLocationContent struct {
	Body    string       `json:"body"`    // Required. A description of the location e.g. 'Big Ben, London, UK', or some kind of content description for accessibility e.g. 'location attachment'.
	GeoURI  string       `json:"geo_uri"` // Required. A geo URI representing this location.
	Msgtype string       `json:"msgtype"` // Required. Must be 'm.location'.
	Info    LocationInfo `json:"info,omitempty"`
}

type LocationInfo struct {
	ThumbnailURL  string        `json:"thumbnail_url,omitempty"`  // The URL to the thumbnail of the file. Only present if the thumbnail is unencrypted.
	ThumbnailFile EncryptedFile `json:"thumbnail_file,omitempty"` // Information on the encrypted thumbnail file, as specified in End-to-end encryption. Only present if the thumbnail is encrypted.
	ThumbnailInfo ThumbnailInfo `json:"thumbnail_info,omitempty"` // Metadata about the image referred to in thumbnail_url.
}

// https://matrix.org/docs/spec/client_server/latest#m-video
type MessageVideoContent struct {
	Body    string        `json:"body"`       // Required. A description of the video e.g. 'Gangnam style', or some kind of content description for accessibility e.g. 'video attachment'.
	Info    VideoInfo     `json:",omitempty"` // Metadata about the video clip referred to in url.
	Msgtype string        `json:"msgtype"`    // Required. Must be 'm.video'.
	URL     string        `json:"url"`        // Required. Required if the file is unencrypted. The URL (typically MXC URI) to the video clip.
	File    EncryptedFile `json:"file"`       // Required if the file is encrypted. Information on the encrypted file, as specified in End-to-end encryption.
}

type VideoInfo struct {
	Duration      int           `json:"duration,omitempty"`       // The duration of the video in milliseconds.
	H             int           `json:"h,omitempty"`              // The height of the video in pixels.
	W             int           `json:"w,omitempty"`              // The width of the video in pixels.
	Mimetype      string        `json:"mimetype,omitempty"`       // The mimetype of the video e.g. video/mp4.
	Size          int           `json:"size,omitempty"`           // The size of the video in bytes.
	ThumbnailURL  string        `json:"thumbnail_url,omitempty"`  // The URL (typically MXC URI) to an image thumbnail of the video clip. Only present if the thumbnail is unencrypted.
	ThumbnailFile EncryptedFile `json:"thumbnail_file,omitempty"` // Information on the encrypted thumbnail file, as specified in End-to-end encryption. Only present if the thumbnail is encrypted.
	ThumbnailInfo ThumbnailInfo `json:"thumbnail_info,omitempty"` // Metadata about the image referred to in thumbnail_url.
}

// https://matrix.org/docs/spec/client_server/latest#extensions-to-m-message-msgtypes
type EncryptedFile struct {
	URL    string            `json:"url"`    // Required. The URL to the file.
	Key    JWK               `json:"key"`    // Required. A JSON Web Key object.
	Iv     string            `json:"iv"`     // Required. The Initialisation Vector used by AES-CTR, encoded as unpadded base64.
	Hashes map[string]string `json:"hashes"` // Required. A map from an algorithm name to a hash of the ciphertext, encoded as unpadded base64. Clients should support the SHA-256 hash, which uses the key sha256.
	V      string            `json:"v"`      // Required. Version of the encrypted attachments protocol. Must be v2.
}

type JWK struct {
	Kty    string   `json:"kty"`     // Required. Key type. Must be oct.
	KeyOps []string `json:"key_ops"` // Required. Key operations. Must at least contain encrypt and decrypt.
	Alg    string   `json:"alg"`     // Required. Algorithm. Must be A256CTR.
	K      string   `json:"k"`       // Required. The key, encoded as urlsafe unpadded base64.
	Ext    bool     `json:"ext"`     // Required. Extractable. Must be true. This is a W3C extension.
}
