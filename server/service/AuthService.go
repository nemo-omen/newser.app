package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"newser.app/infra/dto"
	"newser.app/infra/repository"
	"newser.app/model"
)

var (
	defaultUserCollections = []string{"read", "unread", "saved"}
)

type AuthService struct {
	userRepo       repository.UserRepository
	collectionRepo repository.CollectionRepository
}

func NewAuthService(userRepo repository.UserRepository, collectionRepo repository.CollectionRepository) AuthService {
	return AuthService{
		userRepo:       userRepo,
		collectionRepo: collectionRepo,
	}
}

func (s *AuthService) Login(email, password string) (*model.User, error) {
	hashedPassword, err := s.userRepo.GetHashedPasswordByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("there was an error checking the password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("passwordError: does not match")
	}

	u, _ := s.userRepo.FindByEmail(email)
	return u, nil
}

func (s *AuthService) Signup(email, hashedPassword string) (*model.User, error) {
	udto := &dto.UserDTO{
		Email:          email,
		HashedPassword: hashedPassword,
	}
	u, err := s.userRepo.Create(udto)
	if err != nil {
		return nil, err
	}
	for _, title := range defaultUserCollections {
		collection := model.NewCollection(title, u.Id)
		_, err := s.collectionRepo.Create(collection)
		if err != nil {
			fmt.Printf("error creating %v collection: %v", collection, err.Error())
		}
	}
	return u, nil
}

func (s *AuthService) Logout(userId int64) error {
	return nil
}

func (s *AuthService) GetUserById(userId int64) (*model.User, error) {
	u, err := s.userRepo.Get(userId)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *AuthService) GetUserByEmail(email string) (*model.User, error) {
	u, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return u, nil
}
