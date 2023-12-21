package entity

type User struct {
	ID   string
	Username string
	Password string
	Salt string
}

func NewUser(
	id   string,
	username string,
	password string,
	salt string,
) *User {
	return &User{
		ID: id,
		Username: username,
		Password: password,
		Salt: salt,
	}
}