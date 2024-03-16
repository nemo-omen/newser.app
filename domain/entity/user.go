package entity

import (
	"time"
)

type User struct {
	*Person
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email string) (*User, error) {
	person, err := NewPerson(name, email)
	if err != nil {
		return &User{}, err
	}

	return &User{
		Person:    person,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
