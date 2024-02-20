package view

import "github.com/mmcdole/gofeed"

type SearchPageProps struct {
	Error string
	Feeds []*gofeed.Feed
}
