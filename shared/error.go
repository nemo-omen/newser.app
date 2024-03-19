package shared

import (
	"encoding/json"
	"fmt"
)

type AppError struct {
	Err     error
	Msg     string
	Origin  string
	ErrType string
}

func NewAppError(err error, msg, origin, errType string) *AppError {
	fmt.Println("NewAppError: ")
	fmt.Println(err)
	fmt.Println(msg)
	fmt.Println(origin)
	fmt.Println(errType)
	return &AppError{
		Err:     err,
		Msg:     msg,
		Origin:  origin,
		ErrType: errType,
	}
}

func (a AppError) Error() string {
	return a.Err.Error()
}

func (a AppError) String() string {
	j, _ := json.MarshalIndent(a, "", "  ")
	return string(j)
}

func (a AppError) Print() {
	fmt.Println("AppError:")
	fmt.Println(a.Err)
	fmt.Println(a.ErrType)
	fmt.Println(a.Origin)
	fmt.Println(a.Msg)
}
