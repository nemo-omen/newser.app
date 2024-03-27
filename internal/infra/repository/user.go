package repository

import (
	"newser.app/internal/dto"
)

type UserRepository interface {
	Create(user *dto.UserDTO) error
	GetByEmail(email string) (*dto.UserDTO, error)
	GetByID(id string) (*dto.UserDTO, error)
	Update(user *dto.UserDTO) error
	Delete(id string) error
}
