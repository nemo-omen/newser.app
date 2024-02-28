package service

import (
	"newser.app/infra/repository"
	"newser.app/model"
)

type SubscriptionService struct {
	subRepo  repository.SubscriptionRepository
	feedRepo repository.NewsfeedRepository
}

func NewSubscriptionService(
	sr repository.SubscriptionRepository,
	fr repository.NewsfeedRepository,
) SubscriptionService {
	return SubscriptionService{subRepo: sr, feedRepo: fr}
}

func (s SubscriptionService) Subscribe(f *model.Subscription, userId int) (model.Subscription, error) {
	// persist subscription
	// persist feeds & items
	// add items to user's unread collection
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
