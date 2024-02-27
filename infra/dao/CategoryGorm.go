package dao

import "gorm.io/gorm"

type CategoryGorm struct {
	gorm.Model
	Term      string         `gorm:"not null"`
	Articles  []ArticleGorm  `gorm:"many2many:article_categories"`
	Newsfeeds []NewsfeedGorm `gorm:"many2many:newsfeed_categories"`
}

func (CategoryGorm) TableName() string {
	return "categories"
}
