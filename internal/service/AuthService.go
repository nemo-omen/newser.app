package service

import (
	"newser.app/internal/model"
	"newser.app/internal/repository"
)

type AuthService struct {
	UserRepo repository.Repository[model.User, model.UserGorm]
}

func NewAuthService(dsn string) AuthService {
	return AuthService{
		UserRepo: repository.NewUserGormRepo(dsn),
	}
}

func (s AuthService) Login(email, password string) (model.User, error) {
	return model.User{}, nil
}

func (s AuthService) Signup(email, hashedPassword string) (model.User, error) {
	ug := model.UserGorm{
		Email:          email,
		HashedPassword: hashedPassword,
	}
	u, err := s.UserRepo.Create(ug)
	if err != nil {
		return model.User{}, err
	}
	return u, nil
}

func (s AuthService) Logout(userId uint) error {
	return nil
}

func (s AuthService) GetUser(userId uint) (model.User, error) {
	u, err := s.UserRepo.Get(userId)
	if err != nil {
		return u, err
	}
	return u, nil
}
