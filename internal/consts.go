package internal

const (
	Version = "r0.5.0"
)

// https://matrix.org/docs/spec/client_server/latest#authentication-types
type authenticationType string

const (
	M_LOGIN_PASSWORD       authenticationType = "m.login.password"
	M_LOGIN_RECAPTCHA      authenticationType = "m.login.recaptcha"
	M_LOGIN_OAUTH2         authenticationType = "m.login.oauth2"
	M_LOGIN_EMAIL_IDENTITY authenticationType = "m.login.email.identity"
	M_LOGIN_MSISDN         authenticationType = "m.login.msisdn"
	M_LOGIN_TOKEN          authenticationType = "m.login.token"
	M_LOGIN_DUMMY          authenticationType = "m.login.dummy"
)

// https://matrix.org/docs/spec/client_server/latest#identifier-types
type identifierType string

const (
	// https://matrix.org/docs/spec/client_server/latest#matrix-user-id
	M_ID_USER identifierType = "m.id.user"

	// https://matrix.org/docs/spec/client_server/latest#third-party-id
	M_ID_THIRDPARTY identifierType = "m.id.thirdparty"

	// https://matrix.org/docs/spec/client_server/latest#phone-number
	M_ID_PHONE identifierType = "m.id.phone"
)
