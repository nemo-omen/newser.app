package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
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

type NewsfeedGorm struct {
	gorm.Model
	Title           string          `gorm:"not null"`
	SiteUrl         string          `gorm:"not null"`
	FeedUrl         string          `gorm:"unique;not null"`
	Image           *Image          `gorm:"type:json"`
	Slug            string          `gorm:"not null"`
	Articles        []ArticleGorm   `gorm:"not null"`
	Authors         []PersonGorm    `gorm:"many2many:newsfeed_people"`
	Categories      []*CategoryGorm `gorm:"many2many:newsfeed_categories"`
	Published       string          `gorm:"not null"`
	PublishedParsed *time.Time      `gorm:"not null"`
	Description     *string
	Updated         string
	UpdatedParsed   *time.Time
	Copyright       *string
	Language        *string
	FeedType        string
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
