package usecase

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidSite      = errors.New("invalid site")
	ErrInvalidFeed      = errors.New("invalid feed")
	ErrInvalidURL       = errors.New("invalid url")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrPasswordTooShort = errors.New("password too short")
)

type ServiceError struct {
	Fn  string
	Err error
}

type ResponseError struct {
	Code int
	Fn   string
	Err  error
}

type AuthorizationError struct {
	Fn  string
	Err error
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("service error: %s, %v", e.Fn, e.Err)
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("response code error: %s, %v", e.Fn, e.Err)
}

func (e *AuthorizationError) Error() string {
	return fmt.Sprintf("authorization error: %s, %v", e.Fn, e.Err)
}
