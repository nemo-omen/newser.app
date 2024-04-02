package entity

import "time"

type SearchResultItem struct {
	Title           string
	Description     string
	Content         string
	Link            string
	Links           []string
	Updated         string
	UpdatedParsed   *time.Time
	Published       string
	PublishedParsed *time.Time
	Authors         []*SearchResultPerson
	GUID            string
	Image           *SearchResultImage
	Categories      []string
}
