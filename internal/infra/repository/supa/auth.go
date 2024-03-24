package supa

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"
	"newser.app/internal/domain/entity"
	"newser.app/internal/dto"
)

type AuthSupaRepo struct {
	client supa.Client
}

func NewAuthPGRepo(client supa.Client) *AuthSupaRepo {
	return &AuthSupaRepo{client}
}

func (r *AuthSupaRepo) CreateUser(ctx context.Context, userDao dto.UserDAO, collections []*dto.CollectionDTO) error {
	userCreds := supa.UserCredentials{
		Email:    userDao.Email,
		Password: userDao.Password,
	}
	_, err := r.client.Auth.SignUp(ctx, userCreds)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthSupaRepo) SignIn(ctx context.Context, email, password string) (*supa.AuthenticatedDetails, error) {
	userCreds := supa.UserCredentials{
		Email:    email,
		Password: password,
	}
	var ad = &supa.AuthenticatedDetails{}
	ad, err := r.client.Auth.SignIn(ctx, userCreds)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (r *AuthSupaRepo) FindByEmail(ctx context.Context, token string) (*entity.User, error) {
	return nil, nil
}

func (r *AuthSupaRepo) FindByID(id entity.ID) (*entity.User, error) {
	fmt.Println("FindByID not implemented yet.")
	return nil, nil
}

func (r *AuthSupaRepo) GetHashedPasswordByEmail(email string) (string, error) {
	fmt.Println("GetHashedPasswordByEmail not implemented yet.")
	return "", nil
}
