package model

import "newser.app/shared/util"

type Collection struct {
	Id     int64
	Title  string
	UserId int64
	Slug   string
}

func NewCollection(title string, userId int64) Collection {
	return Collection{
		Title:  title,
		UserId: userId,
		Slug:   util.Slugify(title),
	}
}
