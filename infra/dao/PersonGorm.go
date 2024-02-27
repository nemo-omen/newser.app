package dao

import "gorm.io/gorm"

type PersonGorm struct {
	gorm.Model
	Email     string
	Name      string
	Articles  []ArticleGorm  `gorm:"many2many:article_people"`
	Newsfeeds []NewsfeedGorm `gorm:"many2many:newsfeed_people"`
}

func (PersonGorm) TableName() string {
	return "people"
}
