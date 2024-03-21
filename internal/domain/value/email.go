package value

import (
	"net/mail"

	"newser.app/shared"
)

type Email string

func NewEmail(str string) (Email, error) {
	parsed, err := mail.ParseAddress(str)
	if err != nil {
		appErr := shared.NewAppError(
			ErrInvalidInput,
			"Not a valid email address",
			"NewEmail",
			"value.Email",
		)
		return Email(""), appErr
	}
	return Email(parsed.Address), nil
}

func (e Email) String() string {
	return string(e)
}
