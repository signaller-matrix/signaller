package matrix

import (
	"encoding/json"
)

type ApiError struct {
	Code    string `json:"errcode"`
	Message string `json:"error"`
}

var (
	// https://matrix.org/docs/spec/client_server/latest#api-standards

	M_FORBIDDEN      = ApiError{"M_FORBIDDEN", ""}      // Forbidden access, e.g. joining a room without permission, failed login.
	M_UNKNOWN_TOKEN  = ApiError{"M_UNKNOWN_TOKEN", ""}  // The access token specified was not recognised.
	M_MISSING_TOKEN  = ApiError{"M_MISSING_TOKEN", ""}  // No access token was specified for the request.
	M_BAD_JSON       = ApiError{"M_BAD_JSON", ""}       // Request contained valid JSON, but it was malformed in some way, e.g. missing required keys, invalid values for keys.
	M_NOT_JSON       = ApiError{"M_NOT_JSON", ""}       // Request did not contain valid JSON.
	M_NOT_FOUND      = ApiError{"M_NOT_FOUND", ""}      // No resource was found for this request.
	M_LIMIT_EXCEEDED = ApiError{"M_LIMIT_EXCEEDED", ""} // Too many requests have been sent in a short period of time. Wait a while then try again.
	M_UNKNOWN        = ApiError{"M_UNKNOWN", ""}        // An unknown error has occurred.

	M_UNRECOGNIZED                    = ApiError{"M_UNRECOGNIZED", ""}                    // The server did not understand the request.
	M_UNAUTHORIZED                    = ApiError{"M_UNAUTHORIZED", ""}                    // The request was not correctly authorized. Usually due to login failures.
	M_USER_IN_USE                     = ApiError{"M_USER_IN_USE", ""}                     // Encountered when trying to register a user ID which has been taken.
	M_INVALID_USERNAME                = ApiError{"M_INVALID_USERNAME", ""}                // Encountered when trying to register a user ID which is not valid.
	M_ROOM_IN_USE                     = ApiError{"M_ROOM_IN_USE", ""}                     // Sent when the room alias given to the createRoom API is already in use.
	M_INVALID_ROOM_STATE              = ApiError{"M_INVALID_ROOM_STATE", ""}              // Sent when the initial state given to the createRoom API is invalid.
	M_THREEPID_IN_USE                 = ApiError{"M_THREEPID_IN_USE", ""}                 // Sent when a threepid given to an API cannot be used because the same threepid is already in use.
	M_THREEPID_NOT_FOUND              = ApiError{"M_THREEPID_NOT_FOUND", ""}              // Sent when a threepid given to an API cannot be used because no record matching the threepid was found.
	M_THREEPID_AUTH_FAILED            = ApiError{"M_THREEPID_AUTH_FAILED", ""}            // Authentication could not be performed on the third party identifier.
	M_THREEPID_DENIED                 = ApiError{"M_THREEPID_DENIED", ""}                 // The server does not permit this third party identifier. This may happen if the server only permits, for example, email addresses from a particular domain.
	M_SERVER_NOT_TRUSTED              = ApiError{"M_SERVER_NOT_TRUSTED", ""}              // The client's request used a third party server, eg. identity server, that this server does not trust.
	M_UNSUPPORTED_ROOM_VERSION        = ApiError{"M_UNSUPPORTED_ROOM_VERSION", ""}        // The client's request to create a room used a room version that the server does not support.
	M_INCOMPATIBLE_ROOM_VERSION       = ApiError{"M_INCOMPATIBLE_ROOM_VERSION", ""}       // The client attempted to join a room that has a version the server does not support. Inspect the room_version property of the error response for the room's version.
	M_BAD_STATE                       = ApiError{"M_BAD_STATE", ""}                       // The state change requested cannot be performed, such as attempting to unban a user who is not banned.
	M_GUEST_ACCESS_FORBIDDEN          = ApiError{"M_GUEST_ACCESS_FORBIDDEN", ""}          // The room or resource does not permit guests to access it.
	M_CAPTCHA_NEEDED                  = ApiError{"M_CAPTCHA_NEEDED", ""}                  // A Captcha is required to complete the request.
	M_CAPTCHA_INVALID                 = ApiError{"M_CAPTCHA_INVALID", ""}                 // The Captcha provided did not match what was expected.
	M_MISSING_PARAM                   = ApiError{"M_MISSING_PARAM", ""}                   // A required parameter was missing from the request.
	M_INVALID_PARAM                   = ApiError{"M_INVALID_PARAM", ""}                   // A parameter that was specified has the wrong value. For example, the server expected an integer and instead received a string.
	M_TOO_LARGE                       = ApiError{"M_TOO_LARGE", ""}                       // The request or entity was too large.
	M_EXCLUSIVE                       = ApiError{"M_EXCLUSIVE", ""}                       // The resource being requested is reserved by an application service, or the application service making the request has not created the resource.
	M_RESOURCE_LIMIT_EXCEEDED         = ApiError{"M_RESOURCE_LIMIT_EXCEEDED", ""}         // The request cannot be completed because the homeserver has reached a resource limit imposed on it. For example, a homeserver held in a shared hosting environment may reach a resource limit if it starts using too much memory or disk space. The error MUST have an admin_contact field to provide the user receiving the error a place to reach out to. Typically, this error will appear on routes which attempt to modify state (eg: sending messages, account data, etc) and not routes which only read state (eg: /sync, get account data, etc).
	M_CANNOT_LEAVE_SERVER_NOTICE_ROOM = ApiError{"M_CANNOT_LEAVE_SERVER_NOTICE_ROOM", ""} // The user is unable to reject an invite to join the server notices room. See the Server Notices module for more information.
)

func (apiError ApiError) JSON() []byte {
	b, _ := json.Marshal(apiError) // TODO: error handler?
	return b
}
