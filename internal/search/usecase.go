package search

import (
	"newser.app/internal/search/dto"
)

type SearchUsecase interface {
	FindFeedUrls(url string) ([]string, error)
	GetFeed(url string) (*dto.FeedDTO, error)
	GetFeeds(urls []string) ([]*dto.FeedDTO, error)
}
