package subscription

import "newser.app/internal/subscription/dto"

type SearchRepository interface {
	FindFeedLinks(url string) ([]string, error)
	GetFeed(url string) (*dto.SearchResultFeedDTO, error)
	GetFeeds(urls []string) ([]*dto.SearchResultFeedDTO, error)
}

type SubscriptionRepository interface{}
