package dto

import (
	"encoding/json"

	"newser.app/internal/user/entity"
)

type UserCreateDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u UserCreateDTO) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}

func (u UserCreateDTO) String() string {
	return string(u.JSON())
}

func (u UserCreateDTO) ToDomain() entity.UserCreate {
	return entity.UserCreate{
		Name:  u.Name,
		Email: u.Email,
	}
}
