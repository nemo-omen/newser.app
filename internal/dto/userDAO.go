package dto

import (
	"newser.app/internal/domain/value"
)

type UserDAO struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewUserDAO(ID, name, email, password string) (*UserDAO, error) {
	hashedPassword, err := value.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &UserDAO{
		ID:       ID,
		Name:     name,
		Email:    email,
		Password: hashedPassword.String(),
	}, nil
}
