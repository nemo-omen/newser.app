package repository

import (
	"newser.app/internal/dto"
)

type SubscriptionRepository interface {
	Create(subscription *dto.SubscriptionDTO) error
	Delete(subscriptionID string) error
	GetNewsfeed(userID, feedID string) (*dto.NewsfeedDTO, error)
	GetArticle(userId, id string) (*dto.ArticleDTO, error)
	AddArticle(userID string, article *dto.ArticleDTO) error
	GetAllArticles(userID string) ([]*dto.ArticleDTO, error)
	GetAllFeeds(userID string) ([]*dto.NewsfeedDTO, error)
	GetFeedsInfo(userId string) ([]*dto.FeedInfoDTO, error)
	Subscribe(userID string, feed dto.NewsfeedDTO) (*dto.SubscriptionDTO, error)
}
