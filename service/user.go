package service

import (
	"current/infra/repository"
	"current/model"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s UserService) SignUpUser(email, password string) (*model.User, error) {
	u := &model.User{
		Email: email,
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	u.HashedPassword = hashed
	s.Repo.Create(u)

	return u, nil
}
