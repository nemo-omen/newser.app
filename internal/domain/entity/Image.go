package entity

import (
	"newser.app/internal/domain/value"
)

type Image struct {
	ID    value.ID
	Title string
	URL   string
}
