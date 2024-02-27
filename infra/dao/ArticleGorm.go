package dao

import (
	"time"

	"gorm.io/gorm"
	"newser.app/model"
)

type ArticleGorm struct {
	gorm.Model
	Title           string          `gorm:"not null"`
	Description     string          `gorm:"not null"`
	Content         string          `gorm:"not null"`
	Links           []string        `gorm:"not null"`
	Author          []*PersonGorm   `gorm:"many2many:article_people"`
	Published       string          `gorm:"not null"`
	PublishedParsed time.Time       `gorm:"not null"`
	Image           *model.Image    `gorm:"type:json"`
	Categories      []*CategoryGorm `gorm:"many2many:category_articles"`
	Slug            string          `gorm:"not null"`
	GUID            *string
	Updated         *string
	UpdatedParsed   *time.Time
}

func (ArticleGorm) TableName() string {
	return "articles"
}
