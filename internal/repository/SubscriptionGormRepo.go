package repository

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"newser.app/internal/model"
)

type SubscriptionGormRepo struct {
	DB *gorm.DB
}

func NewSubscriptionGormRepo(dsn string) SubscriptionGormRepo {
	db, err := gorm.Open(sqlite.Open(dsn))
	if err != nil {
		log.Fatal(err)
	}
	return SubscriptionGormRepo{DB: db}
}

func (r SubscriptionGormRepo) Get(id uint) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) Create(sg model.SubscriptionGorm) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) All() []model.Subscription {
	return []model.Subscription{}
}

func (r SubscriptionGormRepo) Update(sg model.SubscriptionGorm) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (r SubscriptionGormRepo) Delete(id uint) error {
	return nil
}
