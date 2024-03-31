package entity

import (
	"encoding/json"

	"newser.app/internal/domain/value"
	"newser.app/shared"
)

type Category struct {
	ID   ID         `json:"id"`
	Term value.Term `json:"term"`
}

func NewCategory(term string) *Category {
	validTerm, err := value.NewTerm(term)
	if err != nil {
		valErr, ok := err.(shared.AppError)
		if ok {
			valErr.Print()
		}
		return nil
	}
	return &Category{
		ID:   NewID(),
		Term: validTerm,
	}
}

func (c Category) JSON() []byte {
	j, _ := json.MarshalIndent(c, "", "  ")
	return j
}

func (c Category) String() string {
	return string(c.JSON())
}
