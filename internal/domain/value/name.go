package value

import (
	"fmt"

	"newser.app/shared"
)

type Name string

const minChars = 3
const maxChars = 32

func NewName(str string) (Name, error) {
	if !IsLongerThanMinChars(str, minChars) {
		return Name(""), shared.NewAppError(
			ErrInvalidInput,
			fmt.Sprintf("Name must be at least %d characters", minChars),
			"NewName",
			"value.Name",
		)
	}

	if !IsShorterThanMaxChars(str, maxChars) {
		return Name(""), shared.NewAppError(
			ErrInvalidInput,
			fmt.Sprintf("Whoah, that's too long! %d characters max", maxChars),
			"NewName",
			"value.Name",
		)
	}

	if !IsAllowedChars(str) {
		return Name(""), shared.NewAppError(
			ErrInvalidInput,
			"Only letters, numbers, or _ allowed",
			"NewName",
			"value.Name",
		)
	}

	return Name(str), nil
}

func (n Name) String() string {
	return string(n)
}
