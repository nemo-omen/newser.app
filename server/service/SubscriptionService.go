package service

import (
	"github.com/mmcdole/gofeed"
	"newser.app/infra/repository"
	"newser.app/model"
)

type SubscriptionService struct {
	subRepo        repository.SubscriptionRepository
	feedRepo       repository.NewsfeedRepository
	collectionRepo repository.CollectionRepository
}

func NewSubscriptionService(
	sr repository.SubscriptionRepository,
	fr repository.NewsfeedRepository,
	cr repository.CollectionRepository,
) SubscriptionService {
	return SubscriptionService{subRepo: sr, feedRepo: fr, collectionRepo: cr}
}

func (s SubscriptionService) Subscribe(f *gofeed.Feed, userId int) (model.Subscription, error) {
	// transform gofeed.Feed into model.Newsfeed
	// transform Feed.Items into Newsfeed.Articles
	// persist Newsfeed & Articles
	// persistedFeed, err := s.feedRepo.Create(f)
	// if err != nil {
	// 	return model.Subscription{}, err
	// }
	// persist subscription
	// add items to user's unread collection
	// for _, article := persistedFeed.Articles {
	// 	s.collectionRepo.Create()
	// }
	// return subscription model
	return model.Subscription{}, nil
}

func (s SubscriptionService) Unsubscribe(feedId, userId uint) error {
	// return err if failure only
	return nil
}

func (s SubscriptionService) All(userId int64) ([]model.Subscription, error) {
	ss, err := s.subRepo.All(userId)
	if err != nil {
		return ss, err
	}
	return ss, nil
}

func (s SubscriptionService) Get(subscriptionId uint) (model.Subscription, error) {
	return model.Subscription{}, nil
}
