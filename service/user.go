package service

import (
	"current/domain"
	"current/infra/repository"

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

func (s UserService) SignUpUser(email, password string) (*domain.User, error) {
	u := &domain.User{
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
