package matrix

type Membership string

const (
	MembershipInvite Membership = "invite"
	MembershipJoin   Membership = "join"
	MembershipKnock  Membership = "knock"
	MembershipLeave  Membership = "leave"
	MembershipBan    Membership = "ban"
)

type SetPresence string

const (
	SetPresenceOffline     SetPresence = "offline"
	SetPresenceOnline      SetPresence = "online"
	SetPresenceUnavailable SetPresence = "unavailable"
)

type JoinRule string

const (
	JoinRulePublic  JoinRule = "public"
	JoinRuleKnock   JoinRule = "knock"
	JoinRuleInvite  JoinRule = "invite"
	JoinRulePrivate JoinRule = "private"
)

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

// https://matrix.org/docs/spec/client_server/r0.4.0.html#id207
type IdentifierType string

const (
	IdentifierTypeUser       IdentifierType = "m.id.user"       // https://matrix.org/docs/spec/client_server/r0.4.0.html#id208
	IdentifierTypeThirdparty IdentifierType = "m.id.thirdparty" // https://matrix.org/docs/spec/client_server/r0.4.0.html#id209
	IdentifierTypePhone      IdentifierType = "m.id.phone"      // https://matrix.org/docs/spec/client_server/r0.4.0.html#id210
)

// Authentication types
// https://matrix.org/docs/spec/client_server/r0.4.0.html#id198
type AuthenticationType string

const (
	// Password-based
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id199
	AuthenticationTypePassword AuthenticationType = "m.login.password"

	// Google ReCaptcha
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id200
	AuthenticationTypeRecaptcha AuthenticationType = "m.login.recaptcha"

	// OAuth2-based
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id202
	AuthenticationTypeOauth2 AuthenticationType = "m.login.oauth2"

	// Email-based (identity server)
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id203
	AuthenticationTypeEmail AuthenticationType = "m.login.email.identity"

	// Token-based
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id201
	AuthenticationTypeToken AuthenticationType = "m.login.token"

	// Dummy Auth
	// https://matrix.org/docs/spec/client_server/r0.4.0.html#id204
	AuthenticationTypeDummy AuthenticationType = "m.login.dummy"
)

type VisibilityType string

const (
	VisibilityTypePrivate = "private"
	VisibilityTypePublic  = "public"
)
