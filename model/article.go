package model

import (
	"time"
)

type Article struct {
	ID              uint
	Title           string
	Description     string
	Content         string
	Links           []string
	Authors         []*Person
	Published       string
	PublishedParsed time.Time
	Updated         string
	UpdatedParsed   time.Time
	Image           *Image
	Categories      []string
	GUID            string
	Slug            string
	FeedId          uint
	FeedTitle       string
	FeedUrl         string
}
