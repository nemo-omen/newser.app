package repository

import (
	"newser.app/internal/domain/entity"
	"newser.app/internal/dto"
)

type AuthRepository interface {
	CreateUser(userDao dto.UserDAO, collections []*entity.Collection) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id entity.ID) (*entity.User, error)
	GetHashedPasswordByEmail(email string) (string, error)
}
