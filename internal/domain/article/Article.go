package article

import (
	"newser.app/internal/domain/entity"
	"newser.app/internal/domain/value"
)

type Article struct {
	item    *entity.Item
	author  *entity.Person
	authors []*entity.Person
	image   *entity.Image
	notes   []*entity.Note
	feed    value.ID
	slug    string
}
