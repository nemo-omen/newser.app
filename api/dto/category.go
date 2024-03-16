package dto

import "newser.app/domain/entity"

type CategoryDTO struct {
	ID   string `json:"id"`
	Term string `json:"term"`
}

func (d CategoryDTO) FromDomain(e *entity.Category) CategoryDTO {
	return CategoryDTO{
		ID:   e.ID.String(),
		Term: e.Term.String(),
	}
}
