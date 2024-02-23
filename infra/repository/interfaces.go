package repository

import (
	"current/model"
)

type UserRepository interface {
	Create(u *model.User) (*model.User, error)
	Get(id uint) (*model.User, error)
	Update(u *model.User) (*model.User, error)
	Delete(id uint)
}
