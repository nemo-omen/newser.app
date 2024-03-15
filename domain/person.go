package domain

import (
	"github.com/google/uuid"
	"newser.app/domain/value"
)

type PersonDTO struct {
	ID    uuid.UUID
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

type PersonDAO struct {
	ID    uuid.UUID `db:"id"`
	Name  string    `db:"name"`
	Email string    `db:"email"`
}

type Person struct {
	ID        uuid.UUID
	Name      value.Name
	Email     value.Email
	Articles  []uuid.UUID
	Newsfeeds []uuid.UUID
}

func NewPerson(name, email string) (Person, error) {
	validName, err := value.NewName(name)
	if err != nil {
		return Person{}, ErrInvalidInput
	}
	validEmail, err := value.NewEmail(email)
	if err != nil {
		return Person{}, ErrInvalidInput
	}

	return Person{
		ID:    uuid.New(),
		Name:  validName,
		Email: validEmail,
	}, nil
}
