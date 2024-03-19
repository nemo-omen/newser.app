package repository

import (
	"newser.app/internal/dto"
)

type SubscriptionRepository interface {
	Create(subscription *dto.SubscriptionDTO) error
	Delete(subscriptionID string) error
	GetAllArticles(userID string) ([]*dto.ArticleDTO, error)
	GetAllFeeds(userID string) ([]*dto.NewsfeedDTO, error)
	GetFeedsInfo(feedID string) (*dto.FeedInfoDTO, error)
	Subscribe(userID string, feed dto.NewsfeedDTO) error
}
