package subscription

import "github.com/mmcdole/gofeed"

type SearchUsecase interface {
	FindFeedUrls(url string) ([]string, error)
	GetFeed(url string) (*gofeed.Feed, error)
	GetFeeds(urls []string) ([]*gofeed.Feed, error)
}
