package value

import (
	"fmt"

	"newser.app/shared"
)

type Term string

func NewTerm(term string) (Term, error) {
	if !IsLongerThanMinChars(term, 3) {
		return Term(""), shared.NewAppError(
			ErrInvalidInput,
			fmt.Sprintf("Must be at least %d characters", minChars),
			"NewTerm",
			"value.Term",
		)
	}

	if !IsShorterThanMaxChars(term, 32) {
		return Term(""), shared.NewAppError(
			ErrInvalidInput,
			fmt.Sprintf("Whoah, that's too long! %d characters max", maxChars),
			"NewTerm",
			"value.Term",
		)
	}

	if !IsAllowedChars(term) {
		return Term(""), shared.NewAppError(
			ErrInvalidInput,
			"Only letters, numbers, or _ allowed",
			"NewTerm",
			"value.Term",
		)
	}

	return Term(term), nil
}

func (t Term) String() string {
	return string(t)
}
