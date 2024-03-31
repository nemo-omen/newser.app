package dto

import (
	"encoding/json"
	"time"

	"newser.app/internal/domain/entity"
)

type UserDTO struct {
	PersonDTO
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u UserDTO) FromDomain(user *entity.User) UserDTO {
	return UserDTO{
		PersonDTO: PersonDTO{}.FromDomain(user.Person),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (u UserDTO) FromJSON(data []byte) (UserDTO, error) {
	err := json.Unmarshal(data, &u)
	return u, err
}

func (u UserDTO) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}
