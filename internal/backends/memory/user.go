package memory

type User struct {
	name     string
	password string
	Tokens   map[string]Token

	backend *Backend
}

func (user *User) ID() string {
	return "@" + user.name + ":" + user.backend.hostname
}

func (user *User) Name() string {
	return user.name
}

func (user *User) Password() string {
	return user.password
}
