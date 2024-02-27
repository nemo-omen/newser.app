package repository

import (
	"newser.app/infra/dto"
	"newser.app/model"
)

type UserRepository interface {
	Get(udto dto.UserDTO) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindById(id uint) (model.User, error)
	GetHashedPasswordByEmail(email string) (string, error)
	All() []model.User
	Create(udto dto.UserDTO) (model.User, error)
	Update(udto dto.UserDTO) (model.User, error)
	Delete(id uint) error
}

type NewsfeedRepository interface {
	Get(id uint) (model.Newsfeed, error)
	Create(n model.Newsfeed) (model.Newsfeed, error)
	Update(n model.Newsfeed) (model.Newsfeed, error)
	Delete(id uint) error
	FindBySlug(slug string) (model.Newsfeed, error)
}

type SubscriptionRepository interface {
	Get(id uint) (model.Subscription, error)
	Create(model.Subscription) (model.Subscription, error)
	All(userId uint) ([]model.Subscription, error)
	Update(model.Subscription) (model.Subscription, error)
	Delete(id uint) error
	FindBySlug(slug string) (model.Subscription, error)
}
