package dao

import (
	"time"

	"gorm.io/gorm"
	"newser.app/model"
)

type ArticleGorm struct {
	gorm.Model
	Title           string
	Description     string
	Content         string
	Links           []string
	Author          []PersonGorm `gorm:"many2many:article_people"`
	Published       string
	PublishedParsed time.Time
	Image           *model.Image   `gorm:"type:json"`
	Categories      []CategoryGorm `gorm:"many2many:category_articles"`
	Slug            string
	GUID            *string
	Updated         *string
	UpdatedParsed   *time.Time
}

func (ArticleGorm) TableName() string {
	return "articles"
}
