package dto

import "newser.app/internal/search/entity"

type PersonDTO struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (s *PersonDTO) ToDomain() *entity.Person {
	if s == nil {
		return nil
	}

	return &entity.Person{
		Name:  s.Name,
		Email: s.Email,
	}
}

func PersonToDTO(s *entity.Person) *PersonDTO {
	if s == nil {
		return nil
	}

	return &PersonDTO{
		Name:  s.Name,
		Email: s.Email,
	}
}
