package dto

import (
	"encoding/json"

	"newser.app/domain/entity"
)

type PersonDTO struct {
	ID        entity.ID   `json:"id,omitempty"`
	Name      string      `json:"name,omitempty"`
	Email     string      `json:"email,omitempty"`
	Articles  []entity.ID `json:"articles,omitempty"`
	Newsfeeds []entity.ID `json:"newsfeeds,omitempty"`
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

func (p PersonDTO) FromDomain(person *entity.Person) PersonDTO {
	return PersonDTO{
		ID:    person.ID,
		Name:  person.Name.String(),
		Email: person.Email.String(),
	}
}
