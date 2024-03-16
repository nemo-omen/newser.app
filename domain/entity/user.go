package entity

import (
	"encoding/json"
	"time"
)

type User struct {
	*Person   `json:"person,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
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

func (u User) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}

func (u User) String() string {
	return string(u.JSON())
}
