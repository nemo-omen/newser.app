package repository

import (
	"newser.app/internal/dto"
)

type ArticleRepository interface {
	Get(articleID string) (*dto.ArticleDTO, error)
	GetBySlug(slug string) (*dto.ArticleDTO, error)
	GetByURL(url string) (*dto.ArticleDTO, error)
	//? AddAnnotation(articleID, annotationID string) error
	//? RemoveAnnotation(articleID, annotationID string) error
}
