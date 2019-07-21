package memory

type User struct {
	id       string
	name     string
	password string
	Tokens   map[string]Token

	backend *Backend
}

func (user *User) ID() string {
	return "@" + user.id + ":" + user.backend.hostname
}

func (user *User) Name() string {
	return user.name
}

func (user *User) Password() string {
	return user.password
}
