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
	Get(interface{}) (interface{}, error)
	Create(interface{}) (interface{}, error)
	Update(interface{}) (interface{}, error)
	Delete(id uint) error
}

type SubscriptionRepository interface{}
