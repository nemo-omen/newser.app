package model

import (
	"time"

	"github.com/mmcdole/gofeed"
	"newser.app/shared/util"
)

type Article struct {
	ID              int64
	Title           string
	Description     string
	Content         string
	ArticleLink     string
	Author          string
	Published       string
	PublishedParsed time.Time
	Updated         string
	UpdatedParsed   time.Time
	Image           string
	Categories      []string
	GUID            string
	Slug            string
	FeedId          int64
	FeedTitle       string
	FeedUrl         string
}

func ArticleFromRemote(ri *gofeed.Item, feedId int64, feedTitle, feedUrl string) Article {
	return Article{
		Title:           ri.Title,
		Description:     ri.Description,
		Content:         ri.Content,
		ArticleLink:     ri.Link,
		Author:          ri.Author.Name,
		Published:       ri.Published,
		PublishedParsed: *ri.PublishedParsed,
		Updated:         ri.Updated,
		UpdatedParsed:   *ri.UpdatedParsed,
		Image:           ri.Image.URL,
		Categories:      ri.Categories,
		GUID:            ri.GUID,
		Slug:            util.Slugify(ri.Title),
		FeedId:          feedId,
		FeedTitle:       feedTitle,
		FeedUrl:         feedUrl,
	}
}
