package entity

import "time"

type SearchResultFeed struct {
	Title           string
	Description     string
	Link            string
	FeedLink        string
	Links           []string
	Updated         string
	UpdatedParsed   *time.Time
	Published       string
	PublishedParsed *time.Time
	Authors         []*SearchResultPerson
	Language        string
	Image           *SearchResultImage
	Copyright       string
	Categories      []string
	Items           []*SearchResultItem
	FeedType        string
}
