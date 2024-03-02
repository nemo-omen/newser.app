package service

import "errors"

var (
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFailure = errors.New("internal failure")
	ErrNotFound        = errors.New("not found")
)
