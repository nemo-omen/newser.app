package repository

import (
	"newser.app/internal/dto"
)

type NewsfeedRepository interface {
	GetNewsfeed(email string) (*dto.NewsfeedDTO, error)
	GetArticle(id string) (*dto.ArticleDTO, error)
}
