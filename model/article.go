package model

import (
	"time"

	"github.com/mmcdole/gofeed"
	"newser.app/shared/util"
)

type Article struct {
	ID              int64     `db:"id"`
	Title           string    `db:"title"`
	Description     string    `db:"description"`
	Content         string    `db:"content"`
	ArticleLink     string    `db:"article_link"`
	Author          string    `db:"author"`
	Published       string    `db:"published"`
	PublishedParsed time.Time `db:"published_parsed"`
	Updated         string    `db:"updated"`
	UpdatedParsed   time.Time `db:"updated_parsed"`
	Image           string    `db:"image"`
	Categories      []string  `db:"categories"`
	GUID            string    `db:"guid"`
	Slug            string    `db:"slug"`
	FeedId          int64     `db:"feed_id"`
	FeedTitle       string    `db:"feed_title"`
	FeedUrl         string    `db:"feed_url"`
}

func ArticleFromRemote(ri *gofeed.Item, feedId int64, feedTitle, feedUrl string) Article {
	var imgUrl string
	if ri.Image != nil {
		imgUrl = ri.Image.URL
	}

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
		Image:           imgUrl,
		Categories:      ri.Categories,
		GUID:            ri.GUID,
		Slug:            util.Slugify(ri.Title),
		FeedId:          feedId,
		FeedTitle:       feedTitle,
		FeedUrl:         feedUrl,
	}
}
