package search

import (
	"newser.app/internal/search/entity"
)

type SearchRepository interface {
	FindFeedLinks(url string) ([]string, error)
	GetFeed(url string) (*entity.Feed, error)
	GetFeeds(urls []string) ([]*entity.Feed, error)
}
