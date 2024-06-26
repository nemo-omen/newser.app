package dto

import "newser.app/internal/domain/entity"

type CategoryDTO struct {
	ID   string `json:"id" db:"id"`
	Term string `json:"term" db:"term"`
}

func (d CategoryDTO) FromDomain(c *entity.Category) CategoryDTO {
	if c == nil {
		return CategoryDTO{}
	}
	return CategoryDTO{
		ID:   c.ID.String(),
		Term: c.Term.String(),
	}
}
