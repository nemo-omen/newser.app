package value

import (
	"encoding/json"
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

func (v ValueError) Unwrap() error {
	return v.Err
}

func (v ValueError) String() string {
	j, _ := json.MarshalIndent(v, "", "  ")
	return string(j)
}
