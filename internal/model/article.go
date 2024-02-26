package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID              uint
	Title           string
	Description     string
	Content         string
	Links           []string
	Authors         []*Person
	Published       string
	PublishedParsed time.Time
	Updated         string
	UpdatedParsed   time.Time
	Image           *Image
	Categories      []string
	GUID            string
	Slug            string
	FeedId          uint
	FeedTitle       string
	FeedUrl         string
}

type ArticleGorm struct {
	gorm.Model
	Title           string          `gorm:"not null"`
	Description     string          `gorm:"not null"`
	Content         string          `gorm:"not null"`
	Links           []string        `gorm:"not null"`
	Author          []*PersonGorm   `gorm:"many2many:article_people"`
	Published       string          `gorm:"not null"`
	PublishedParsed time.Time       `gorm:"not null"`
	Image           *Image          `gorm:"type:json"`
	Categories      []*CategoryGorm `gorm:"many2many:category_articles"`
	Slug            string          `gorm:"not null"`
	GUID            *string
	Updated         *string
	UpdatedParsed   *time.Time
}
