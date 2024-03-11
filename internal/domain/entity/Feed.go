package entity

import (
	"newser.app/internal/domain/value"
)

type Feed struct {
	ID          value.ID
	Title       string
	SiteUrl     string
	FeedUrl     string
	Description string
	Copyright   string
	Language    string
	FeedType    string
}
