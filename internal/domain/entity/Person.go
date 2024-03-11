package entity

import (
	"newser.app/internal/domain/value"
)

type Person struct {
	ID    value.ID
	Name  string
	Email string
}
