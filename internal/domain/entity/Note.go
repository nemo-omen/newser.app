package entity

import (
	"newser.app/internal/domain/value"
)

type Note struct {
	ID      value.ID
	Title   string
	Content string
}
