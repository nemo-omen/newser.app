package model

import (
	"encoding/json"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
)

type Article struct {
	ID              int64     `db:"id"`
	Title           string    `json:"title,omitempty" db:"title"`
	Description     string    `json:"description,omitempty" db:"description"`
	Content         string    `json:"content,omitempty" db:"content"`
	ArticleLink     string    `json:"link,omitempty" db:"article_link"`
	Author          *Person   `json:"author,omitempty" db:"-"`
	Published       string    `json:"published,omitempty" db:"published"`
	PublishedParsed time.Time `json:"publishedParsed,omitempty" db:"published_parsed"`
	Updated         string    `json:"updated,omitempty" db:"updated"`
	UpdatedParsed   time.Time `json:"updatedParsed,omitempty" db:"updated_parsed"`
	Image           *Image    `json:"image,omitempty" db:"-"`
	Categories      []string  `json:"categories,omitempty" db:"categories"`
	GUID            string    `json:"guid,omitempty" db:"guid"`
	Slug            string    `json:"-" db:"slug"`
	FeedId          int64     `json:"-" db:"feed_id"`
	Read            bool      `json:"-" db:"read"`
}

func ArticleFromRemote(ri *gofeed.Item) (*Article, error) {
	a := new(Article)
	p := bluemonday.UGCPolicy()
	m, err := json.Marshal(ri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(m, a)
	if err != nil {
		return nil, err
	}
	a.Content = p.Sanitize(a.Content)

	if a.Description == "" {
		d := strip.StripTags(a.Content)
		a.Description = d
	}

	if len(a.Description) > 87 {
		a.Description = a.Description[:87] + "..."
	}

	return a, nil
}
