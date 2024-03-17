package shared

type AppError struct {
	Err     error
	Msg     string
	Origin  string
	ErrType string
}

func NewAppError(err error, msg, origin, errType string) *AppError {
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
