package dto

import (
	"encoding/json"

	"newser.app/internal/user/entity"
)

type UserIdentityDTO struct {
	ID    int64  `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
}

func (u UserIdentityDTO) JSON() []byte {
	j, _ := json.MarshalIndent(u, "", "  ")
	return j
}

func (u UserIdentityDTO) String() string {
	return string(u.JSON())
}

func (u UserIdentityDTO) ToDomain() entity.UserIdentity {
	return entity.UserIdentity{
		ID:    u.ID,
		Email: u.Email,
	}
}
