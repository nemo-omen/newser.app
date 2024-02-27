package service

import (
	"newser.app/infra/repository"
	"newser.app/model"
)

type SubscriptionService struct {
	Repo repository.SubscriptionRepository
}

func NewSubscriptionService(sr repository.SubscriptionRepository) SubscriptionService {
	return SubscriptionService{
		Repo: sr,
	}
}

func (s SubscriptionService) Subscribe(f *model.Subscription, userId int) (model.Subscription, error) {
	// persist feeds, create subscription
	// return subscription id
	return model.Subscription{}, nil
}

func (s SubscriptionService) Unsubscribe(feedId, userId uint) error {
	// return err if failure only
	return nil
}

func (s SubscriptionService) All() ([]model.Subscription, error) {
	return []model.Subscription{}, nil
}

func (s SubscriptionService) Get(subscriptionId uint) (model.Subscription, error) {
	return model.Subscription{}, nil
}
