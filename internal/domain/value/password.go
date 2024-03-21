package value

import (
	"golang.org/x/crypto/bcrypt"
	"newser.app/shared"
)

type Password string

// NewPassword creates a new hashed Password value object
// and returns an error if the password is too short
// or if the bcrypt hashing fails
func NewPassword(password string) (Password, error) {
	if !IsLongerThanMinChars(password, 6) {
		err := shared.NewAppError(
			ErrInvalidInput,
			"Password must be at least 6 characters",
			"NewPassword",
			"value.Password",
		)
		return Password(""), err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		appErr := shared.NewAppError(
			err,
			"Failed to hash password",
			"NewPassword",
			"value.Password",
		)
		return Password(""), appErr
	}

	return Password(string(bytes)), nil
}

// Compare compares a Password value object with a string
// and returns an error if the comparison fails
func (p Password) Compare(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(password))
}

// String returns the string representation of a Password value object
func (p Password) String() string {
	return string(p)
}
