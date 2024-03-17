package value

import (
	"errors"
)

var (
	ErrInvalidInput     = errors.New("invalid input")
	ErrPasswordTooShort = errors.New("password too short")
)

type ValueError struct {
	Err       error
	Msg       string
	ValueType string
}

func NewValueError(err error, msg, errType string) *ValueError {
	return &ValueError{
		Err:       err,
		Msg:       msg,
		ValueType: errType,
	}
}

func (v ValueError) Error() string {
	return v.Err.Error()
}
