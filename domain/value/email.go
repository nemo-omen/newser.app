package value

import (
	"net/mail"
)

type Email string

func NewEmail(str string) (Email, error) {
	parsed, err := mail.ParseAddress(str)
	if err != nil {
		return Email(""), ErrInvalidInput
	}
	return Email(parsed.Address), nil
}

func (e Email) String() string {
	return string(e)
}
