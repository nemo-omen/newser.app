package repository

import (
	"current/domain"
)

type UserRepository interface {
	Create(u *domain.User) (*domain.User, error)
	Get(id uint) (*domain.User, error)
	Update(u *domain.User) (*domain.User, error)
	Delete(id uint)
}
