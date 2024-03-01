package model

import "newser.app/shared/util"

type Collection struct {
	Id     int64  `db:"id"`
	Title  string `db:"title"`
	UserId int64  `db:"user_ud"`
	Slug   string `db:"slug"`
}

func NewCollection(title string, userId int64) *Collection {
	return &Collection{
		Title:  title,
		UserId: userId,
		Slug:   util.Slugify(title),
	}
}
