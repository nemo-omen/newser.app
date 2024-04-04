package dto

import (
	"encoding/json"

	"newser.app/internal/subscription/entity"
)

type PersonDTO struct {
	ID    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (p PersonDTO) JSON() []byte {
	json, _ := json.Marshal(p)
	return json
}

func (p PersonDTO) FromJSON(data []byte) (PersonDTO, error) {
	err := json.Unmarshal(data, &p)
	return p, err
}

func (p PersonDTO) String() string {
	return string(p.JSON())
}

func (p PersonDTO) ToDomain() entity.Person {
	return entity.Person{
		ID:    p.ID,
		Name:  p.Name,
		Email: p.Email,
	}
}
