package repository

import (
	"gorm.io/gorm"
	"newser.app/infra/dao"
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

func (r SubscriptionGormRepo) Create(sg dao.SubscriptionGorm) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) All() []model.Subscription {
	return []model.Subscription{}
}

func (r SubscriptionGormRepo) Update(sg dao.SubscriptionGorm) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) Delete(id uint) error {
	return nil
}
