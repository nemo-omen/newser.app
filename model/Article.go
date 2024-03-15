package model

import (
	"encoding/json"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"newser.app/shared/util"
)

type Article struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	Title       string `json:"title,omitempty" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	Content     string `json:"content,omitempty" db:"content"`
	ArticleLink string `json:"link,omitempty" db:"article_link"`
	// Author          *Person   `json:"author,omitempty"`
	Person
	Published       string    `json:"published,omitempty" db:"published"`
	PublishedParsed time.Time `json:"publishedParsed,omitempty" db:"published_parsed"`
	Updated         string    `json:"updated,omitempty" db:"updated"`
	UpdatedParsed   time.Time `json:"updatedParsed,omitempty" db:"updated_parsed"`
	Image           *Image    `json:"image,omitempty"`
	Categories      []string  `json:"categories,omitempty" db:"categories"`
	GUID            string    `json:"guid,omitempty" db:"guid"`
	Slug            string    `json:"-" db:"slug"`
	FeedId          int64     `json:"-" db:"feed_id"`
	Read            bool      `json:"-" db:"read"`
	Saved           bool      `json:"-" db:"-"`
	FeedTitle       string    `db:"feed_title" json:"-"`
	FeedUrl         string    `db:"feed_url" json:"-"`
	FeedSiteUrl     string    `db:"feed_site_url" json:"-"`
	FeedSlug        string    `db:"feed_slug" json:"-"`
	FeedImageUrl    string    `db:"feed_image_url" json:"-"`
	FeedImageTitle  string    `db:"feed_image_title" json:"-"`
}

func ArticleFromRemote(ri *gofeed.Item) (*Article, error) {
	a := new(Article)
	p := bluemonday.UGCPolicy()
	// p := bluemonday.NewPolicy()
	p.AllowElements("code")
	p.AllowAttrs("class").OnElements("code")
	m, err := json.Marshal(ri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(m, a)
	if err != nil {
		return nil, err
	}
	a.Content = p.Sanitize(a.Content)

	// rather than using default description,
	// we strip the html from the content
	// and grab a small substring.
	// This is both for consistencyy and for
	// the fact that the lede is usually
	// more interesting than the blurb
	// (in my experience/opinion)
	d := strip.StripTags(a.Content)
	a.Description = d
	a.Slug = util.Slugify(a.Title)

	if len(a.Description) > 240 {
		a.Description = a.Description[:240] + "..."
	}

	return a, nil
}
