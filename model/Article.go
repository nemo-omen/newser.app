package model

import (
	"encoding/json"
	"time"

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
	Slug            string    `db:"slug"`
	FeedId          int64     `db:"feed_id"`
}

func ArticleFromRemote(ri *gofeed.Item) (*Article, error) {
	a := new(Article)
	m, err := json.Marshal(ri)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(m, a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
