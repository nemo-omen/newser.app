package value

import (
	"errors"
	"net/mail"
)

var (
	ErrInvalidValue = errors.New("invalid value")
)

type Email string

func NewEmail(str string) (Email, error) {
	parsed, err := mail.ParseAddress(str)
	if err != nil {
		return Email(""), ErrInvalidValue
	}
	return Email(parsed.Address), nil
}

type Name string

func NewName(str string) (Name, error) {
	if len(str) == 0 {
		return Name(""), ErrInvalidValue
	}
	return Name(str), nil
}
