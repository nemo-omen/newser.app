package entity

import (
	"encoding/json"
	"time"

	"newser.app/internal/domain/value"
)

type User struct {
	Email     value.Email `json:"email,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
}

func NewUser(name, email string) (*User, error) {
	validEmail, err := value.NewEmail(email)

	if err != nil {
		return nil, err
	}

	return &User{
		Email:     validEmail,
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
