package login

type GetLoginReply struct {
	Flows []Flow `json:"flows"` // The homeserver's supported login types
}
