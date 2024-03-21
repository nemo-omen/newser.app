package entity

import (
	"encoding/json"

	"newser.app/internal/domain/value"
	"newser.app/shared"
)

type Person struct {
	ID    ID          `json:"id,omitempty"`
	Name  value.Name  `json:"name,omitempty"`
	Email value.Email `json:"email,omitempty"`
}

func NewPerson(name, email string) (*Person, error) {
	validName, err := value.NewName(name)
	if err != nil {
		valErr, ok := err.(shared.AppError)
		if ok {
			valErr.Print()
		}
		return nil, err
	}
	validEmail, err := value.NewEmail(email)
	if err != nil {
		valErr, ok := err.(shared.AppError)
		if ok {
			valErr.Print()
		}
		return nil, err
	}

	return &Person{
		ID:    NewID(),
		Name:  validName,
		Email: validEmail,
	}, nil
}

func (p Person) JSON() []byte {
	json, _ := json.Marshal(p)
	return json
}

func (p Person) String() string {
	return string(p.JSON())
}
