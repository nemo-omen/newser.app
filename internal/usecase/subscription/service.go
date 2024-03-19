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

func (s *SubscriptionService) Subscribe(userID string, feed gofeed.Feed) error {
	fmt.Printf("Subscribing user %s to feed %s: ", userID, feed.Title)
	feedMapper := mapper.GofeedMapper{}
	newsfeed, err := feedMapper.ToNewsfeed(&feed)
	if err != nil {
		return err
	}
	feedDTO := dto.NewsfeedDTO{}.FromDomain(*newsfeed)
	return s.subscriptionRepo.Subscribe(userID, feedDTO)
}

func (s *SubscriptionService) UnSubscribe(userID, subscriptionID string) error {
	return s.subscriptionRepo.Delete(subscriptionID)
}
