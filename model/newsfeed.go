package model

import (
	"encoding/json"
	"time"
)

type Newsfeed struct {
	ID              uint
	Title           string
	SiteUrl         string // site link
	FeedUrl         string // feed link
	Description     string
	Image           *Image
	Published       string
	PublishedParsed time.Time
	Updated         string
	UpdatedParsed   time.Time
	Copyright       string
	Articles        []*Article
	Author          *Person
	Authors         []*Person
	Language        string
	Categories      []string
	FeedType        string
	Slug            string
}

func (f Newsfeed) String() string {
	json, _ := json.MarshalIndent(f, "", "    ")
	return string(json)
}

func (f Newsfeed) Len() int {
	return len(f.Articles)
}

func (f Newsfeed) Less(i, k int) bool {
	return f.Articles[i].PublishedParsed.Before(
		f.Articles[k].PublishedParsed,
	)
}

func (f Newsfeed) Swap(i, k int) {
	f.Articles[i], f.Articles[k] = f.Articles[k], f.Articles[i]
}
