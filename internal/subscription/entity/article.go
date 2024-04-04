package entity

import (
	"time"
)

type Article struct {
	ID              int64
	Title           string
	Description     string
	Content         string
	Link            string
	Updated         string
	UpdatedParsed   time.Time
	Published       string
	PublishedParsed time.Time
	Author          Person
	GUID            string
	Image           Image
	Categories      []string
	Slug            string
}
