package repository

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrTransactionError = errors.New("transaction error")
	ErrInsertError      = errors.New("insertion error")
	ErrMigrationError   = errors.New("migration error")
)
