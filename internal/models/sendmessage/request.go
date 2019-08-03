package sendmessage

import (
	common "github.com/signaller-matrix/signaller/internal/models/common"
)

type SendMessageRequest struct {
	Body        string             `json:"body"`    // Required. The textual representation of this message.
	MessageType common.MessageType `json:"msgtype"` // Required. The type of message, e.g. m.image, m.text
	RelatesTo   MRelatesTo         `json:"m.relates_to,omitempty"`
}

type MRelatesTo struct {
	InReplyTo MInReplyTo `json:"m.in_reply_to"`
}

type MInReplyTo struct {
	EventID string `json:"event_id"`
}
