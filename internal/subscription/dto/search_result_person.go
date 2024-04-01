package dto

import "newser.app/internal/search/entity"

type SearchResultPersonDTO struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (s *SearchResultPersonDTO) ToDomain() *entity.SearchResultPerson {
	if s == nil {
		return nil
	}

	return &entity.SearchResultPerson{
		Name:  s.Name,
		Email: s.Email,
	}
}
