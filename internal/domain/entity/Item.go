package entity

import (
	"time"

	"newser.app/internal/domain/value"
)

type Item struct {
	ID              value.ID
	Title           string
	Description     string
	Content         string
	Link            string
	Links           []string
	Updated         string
	UpdatedParsed   time.Time
	Published       string
	PublishedParsed time.Time
	GUID            string
}
