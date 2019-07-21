package events

import (
	"encoding/json"
)

// Type is type of event
type Type string

// Event is the basic set of fields all events must have
// https://matrix.org/docs/spec/client_server/latest#event-fields
type Event struct {
	Content json.RawMessage `json:"content"` //	Required. The fields in this object will vary depending on the type of event. When interacting with the REST API, this is the HTTP body.
	Type    Type            `json:"type"`    // Required. The type of event. This SHOULD be namespaced similar to Java package naming conventions e.g. 'com.example.subdomain.event.type'
}
