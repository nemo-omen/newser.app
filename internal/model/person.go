package model

import "gorm.io/gorm"

type Person struct {
	ID       uint
	Name     string
	Email    string
	Articles []uint
}

type PersonGorm struct {
	gorm.Model
	Email     string
	Name      string          `gorm:"not null"`
	Articles  []*ArticleGorm  `gorm:"many2many:article_people"`
	Newsfeeds []*NewsfeedGorm `gorm:"many2many:newsfeed_people"`
}

func (PersonGorm) TableName() string {
	return "people"
}
