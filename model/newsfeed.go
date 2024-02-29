package model

import (
	"encoding/json"
	"time"

	"github.com/mmcdole/gofeed"
	"newser.app/shared/util"
)

type Newsfeed struct {
	ID              int64
	Title           string
	SiteUrl         string // site link
	FeedUrl         string // feed link
	Description     string
	Image           string
	Published       string
	PublishedParsed *time.Time
	Updated         string
	UpdatedParsed   *time.Time
	Copyright       string
	Articles        []Article
	Author          string
	// Authors         []string
	Language   string
	Categories []string
	FeedType   string
	Slug       string
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

func FeedFromRemote(rf gofeed.Feed) Newsfeed {
	return Newsfeed{
		Title:           rf.Title,
		SiteUrl:         rf.Link,
		FeedUrl:         rf.FeedLink,
		Description:     rf.Description,
		Image:           rf.Image.URL,
		Published:       rf.Published,
		PublishedParsed: rf.PublishedParsed,
		Updated:         rf.Updated,
		UpdatedParsed:   rf.UpdatedParsed,
		Copyright:       rf.Copyright,
		Author:          rf.Author.Name,
		Language:        rf.Language,
		Categories:      rf.Categories,
		FeedType:        rf.FeedType,
		Slug:            util.Slugify(rf.Title),
	}
}
