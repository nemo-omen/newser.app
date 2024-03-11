package entity

import (
	"newser.app/internal/domain/value"
)

type Category struct {
	ID   value.ID
	Term string
}
