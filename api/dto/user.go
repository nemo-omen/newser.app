package dto

import (
	"time"
)

type UserDTO struct {
	PersonDTO
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
