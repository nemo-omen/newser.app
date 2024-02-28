package model

import (
	"time"
)

type Article struct {
	ID              int64
	Title           string
	Description     string
	Content         string
	ArticleLink     string
	Authors         []*Person
	Published       string
	PublishedParsed time.Time
	Updated         string
	UpdatedParsed   time.Time
	Image           *Image
	Categories      []string
	GUID            string
	Slug            string
	FeedId          int64
	FeedTitle       string
	FeedUrl         string
}
