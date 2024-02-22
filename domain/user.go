package domain

type User struct {
	Id             uint
	Email          string
	HashedPassword []byte
}

func NewUser(email string) *User {
	u := User{
		Email: email,
	}

	return &u
}
