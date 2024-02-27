package repository

import (
	"gorm.io/gorm"
	"newser.app/model"
)

type SubscriptionGormRepo struct {
	DB *gorm.DB
}

func NewSubscriptionGormRepo(db *gorm.DB) SubscriptionGormRepo {
	return SubscriptionGormRepo{DB: db}
}

func (r SubscriptionGormRepo) Get(id uint) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) Create(s model.Subscription) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) All(userId uint) ([]model.Subscription, error) {
	return []model.Subscription{}, nil
}

func (r SubscriptionGormRepo) Update(s model.Subscription) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) Delete(id uint) error {
	return nil
}

func (r SubscriptionGormRepo) FindBySlug(slug string) (model.Subscription, error) {
	return model.Subscription{}, nil
}
