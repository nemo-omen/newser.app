package model

import (
	"encoding/json"
	"net/url"

	"github.com/mmcdole/gofeed"
	"newser.app/shared/util"
)

type NewsfeedExtended struct {
	Newsfeed
	SubscriptionId int64  `db:"subscription_id"`
	ImageUrl       string `db:"feed_image_url"`
	ImageTitle     string `db:"feed_image_title"`
	UnreadCount    int
}

type Newsfeed struct {
	ID          int64      `db:"id"`
	Title       string     `json:"title,omitempty" db:"title"`
	SiteUrl     string     `json:"link,omitempty" db:"site_url"`
	FeedUrl     string     `json:"feedLink,omitempty" db:"feed_url"`
	Description string     `json:"description,omitempty" db:"description"`
	Image       *Image     `json:"image"`
	Copyright   string     `json:"copyright,omitempty" db:"copyright"`
	Articles    []*Article `json:"articles"`
	Author      *Person    `json:"author"`
	Language    string     `json:"language,omitempty" db:"language"`
	Categories  []string   `json:"categories,omitempty" db:"categories"`
	FeedType    string     `json:"feedType,omitempty" db:"feed_type"`
	Slug        string     `json:"slug" db:"slug"`
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

func FeedFromRemote(rf *gofeed.Feed) (*Newsfeed, error) {
	m, err := json.Marshal(rf)
	nf := new(Newsfeed)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(m, nf)
	if err != nil {
		return nil, err
	}
	nf.Slug = util.Slugify(nf.Title)

	if nf.SiteUrl == "" {
		u, err := url.Parse(nf.FeedUrl)
		if err == nil {
			scheme := u.Scheme
			host := u.Host
			nf.SiteUrl = scheme + host
		}
	}

	return nf, nil
}
