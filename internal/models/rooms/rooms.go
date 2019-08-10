package rooms

type JoinRule string

const (
	Public  JoinRule = "public"
	Knock            = "knock"
	Invite           = "invite"
	Private          = "private"
)
