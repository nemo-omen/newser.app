package repository

import (
	"newser.app/internal/dto"
)

type CollectionRepository interface {
	Create(collection *dto.CollectionDTO) error
	Delete(collectionID string) error
	Get(collectionID string) (*dto.CollectionDTO, error)
	GetByTitle(title, userId string) (*dto.CollectionDTO, error)
	GetBySlug(slug, userId string) (*dto.CollectionDTO, error)
	All(userID string) ([]*dto.CollectionDTO, error)
	AddArticle(collectionID, articleID string) error
	RemoveArticle(collectionID, articleID string) error
	AddNewsfeed(collectionID, newsfeedID string) error
	RemoveNewsfeed(collectionID, newsfeedID string) error
	GetCollectionArticles(collectionID string) ([]*dto.ArticleDTO, error)
	GetCollectionNewsfeeds(collectionID string) ([]*dto.NewsfeedDTO, error)
}
