package entity

import (
	"newser.app/internal/domain/value"
)

type Subscription struct {
	ID    value.ID
	Title string
}
