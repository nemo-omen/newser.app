package newsfeed

import (
	"github.com/google/uuid"
	"newser.app/internal/domain/article"
	"newser.app/internal/domain/entity"
)

type Newsfeed struct {
	ID          uuid.UUID
	title       string
	siteUrl     string
	feedUrl     string
	description string
	image       *entity.Image
	articles    []*article.Article
	author      *entity.Person
	authors     []*entity.Person
	categories  []*entity.Category
	slug        string
}
