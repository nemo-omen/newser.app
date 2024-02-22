package repository

import (
	"current/domain"
	"current/util"
	"fmt"
)

type UserMemRepo struct {
	Users []*domain.User
}

func NewUserMemRepo() *UserMemRepo {
	return &UserMemRepo{
		Users: []*domain.User{},
	}
}

func (r *UserMemRepo) Create(u *domain.User) (*domain.User, error) {
	for _, uu := range r.Users {
		if uu.Email == u.Email {
			return u, fmt.Errorf("email %v already exists", u.Email)
		}
	}

	id := len(r.Users)
	u.Id = uint(id)
	r.Users = append(r.Users, u)

	return u, nil
}

func (r *UserMemRepo) Get(id uint) (*domain.User, error) {
	filtered := util.Filter[*domain.User](r.Users, func(u *domain.User) bool {
		return u.Id == id
	})

	if len(filtered) < 1 {
		return &domain.User{}, nil
	}

	if len(filtered) > 1 {
		return &domain.User{}, fmt.Errorf("found more than one user with id %v", id)
	}

	return filtered[0], nil
}

func (r *UserMemRepo) Update(u *domain.User) (*domain.User, error) {
	storedUser, err := r.Get(u.Id)
	if err != nil {
		return u, err
	}
	storedUser.Email = u.Email
	return storedUser, nil
}

func (r *UserMemRepo) Delete(id uint) {
	ret := make([]*domain.User, 0)

	for _, u := range r.Users {
		if u.Id != id {
			ret = append(ret, u)
		}
	}
	r.Users = ret
}
