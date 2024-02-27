package dao

import (
	"time"

	"gorm.io/gorm"
	"newser.app/model"
)

type NewsfeedGorm struct {
	gorm.Model
	Title           string          `gorm:"not null"`
	SiteUrl         string          `gorm:"not null"`
	FeedUrl         string          `gorm:"unique;not null"`
	Image           *model.Image    `gorm:"type:json"`
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
