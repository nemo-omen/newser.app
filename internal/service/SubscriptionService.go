package service

import (
	"newser.app/internal/model"
	"newser.app/internal/repository"
)

type SubscriptionService struct {
	Repo repository.Repository[model.Subscription, model.SubscriptionGorm]
}

func NewSubscriptionService(dsn string) SubscriptionService {
	return SubscriptionService{
		Repo: repository.NewSubscriptionGormRepo(dsn),
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
