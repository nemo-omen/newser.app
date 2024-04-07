package dto

import (
	"encoding/json"

	"newser.app/internal/user/entity"
)

type UserDetailDTO struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (u UserDetailDTO) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}

func (u UserDetailDTO) String() string {
	return string(u.JSON())
}

func (u UserDetailDTO) ToDomain() entity.UserDetail {
	return entity.UserDetail{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
