package subscription

import (
	"fmt"

	"github.com/mmcdole/gofeed"
	"newser.app/internal/dto"
	"newser.app/internal/infra/mapper"
	"newser.app/internal/infra/repository"
)

type SubscriptionService struct {
	subscriptionRepo repository.SubscriptionRepository
}

func NewSubscriptionService(subscriptionRepo repository.SubscriptionRepository) SubscriptionService {
	return SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

func (s *SubscriptionService) GetAllArticles(userID string) ([]*dto.ArticleDTO, error) {
	return s.subscriptionRepo.GetAllArticles(userID)
}

func (s *SubscriptionService) Subscribe(userID string, feed gofeed.Feed) (*dto.SubscriptionDTO, error) {
	fmt.Printf("Subscribing user %s to feed %s: ", userID, feed.Title)
	feedMapper := mapper.GofeedMapper{}
	newsfeed, err := feedMapper.ToNewsfeed(&feed)
	if err != nil {
		return nil, err
	}
	feedDTO := dto.NewsfeedDTO{}.FromDomain(*newsfeed)
	subscription, err := s.subscriptionRepo.Subscribe(userID, feedDTO)
	return subscription, err
}

func (s *SubscriptionService) UnSubscribe(userID, subscriptionID string) error {
	return s.subscriptionRepo.Delete(subscriptionID)
}

func (s *SubscriptionService) GetNewsfeed(userID, feedID string) (*dto.NewsfeedDTO, error) {
	return s.subscriptionRepo.GetNewsfeed(userID, feedID)
}

func (s *SubscriptionService) GetArticle(articleId string) (*dto.ArticleDTO, error) {
	return s.subscriptionRepo.GetArticle(articleId)
}

func (s *SubscriptionService) GetSidebarLinks(userID string) ([]*dto.FeedInfoDTO, error) {
	return s.subscriptionRepo.GetFeedsInfo(userID)
}
