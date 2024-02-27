package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"newser.app/infra/dto"
	"newser.app/infra/repository"
	"newser.app/model"
)

type AuthService struct {
	UserRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return AuthService{
		UserRepo: userRepo,
	}
}

func (s AuthService) Login(email, password string) (model.User, error) {
	hashedPassword, err := s.UserRepo.GetHashedPasswordByEmail(email)
	if err != nil {
		return model.User{}, fmt.Errorf("there was an error checking the password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return model.User{}, fmt.Errorf("passwordError: does not match")
	}

	u, _ := s.UserRepo.FindByEmail(email)
	return u, nil
}

func (s AuthService) Signup(email, hashedPassword string) (model.User, error) {
	udto := dto.UserDTO{
		Email:          email,
		HashedPassword: hashedPassword,
	}
	u, err := s.UserRepo.Create(udto)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (s AuthService) Logout(userId uint) error {
	return nil
}

func (s AuthService) GetUserById(userId uint) (model.User, error) {
	u, err := s.UserRepo.FindById(userId)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (s AuthService) GetUserByEmail(email string) (model.User, error) {
	u, err := s.UserRepo.FindByEmail(email)
	if err != nil {
		return u, err
	}
	return u, nil
}
