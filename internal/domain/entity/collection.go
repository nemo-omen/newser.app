package entity

import (
	"newser.app/internal/domain/value"
	"newser.app/shared/util"
)

type Collection struct {
	ID        ID
	Title     value.Name
	Slug      string
	UserID    ID
	Articles  []*ID
	Newsfeeds []*ID
}

func NewCollection(title, userId string) (*Collection, error) {
	validName, err := value.NewName(title)
	if err != nil {
		return nil, err
	}

	uID, err := NewIDFromString(userId)
	if err != nil {
		return nil, err
	}

	slug := util.Slugify(title)
	c := &Collection{
		ID:     NewID(),
		Title:  validName,
		Slug:   slug,
		UserID: uID,
	}

	return c, nil
}
