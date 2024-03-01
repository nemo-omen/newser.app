package model

import (
	"encoding/json"
	"time"

	"github.com/mmcdole/gofeed"
	"newser.app/shared/util"
)

type Newsfeed struct {
	ID              int64      `db:"id"`
	Title           string     `db:"title"`
	SiteUrl         string     `db:"site_url"`
	FeedUrl         string     `db:"feed_url"`
	Description     string     `db:"description"`
	Image           string     `db:"image"`
	Published       string     `db:"published"`
	PublishedParsed *time.Time `db:"published_parsed"`
	Updated         string     `db:"updated"`
	UpdatedParsed   *time.Time `db:"updated_parsed"`
	Copyright       string     `db:"copyright"`
	Articles        []*Article `db:"articles"`
	Author          string     `db:"author"`
	Language        string     `db:"language"`
	Categories      []string   `db:"categories"`
	FeedType        string     `db:"feed_type"`
	Slug            string     `db:"slug"`
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

func FeedFromRemote(rf *gofeed.Feed) *Newsfeed {
	return &Newsfeed{
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
