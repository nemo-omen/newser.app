package entity

import "newser.app/domain/value"

type Newsfeed struct {
	ID          ID
	Title       string
	Description string
	FeedLink    value.Link
	SiteLink    value.Link
	Author      *Person
	Language    string
	Image       *Image
	Copyright   string
	Articles    []*Article
}
