package model

import (
	"time"

	"github.com/mmcdole/gofeed"
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

func ArticleFromRemote(ri *gofeed.Item) Article {
	return Article{}
}
