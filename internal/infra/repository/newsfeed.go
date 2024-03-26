package repository

import "newser.app/internal/dto"

type NewsfeedRepository interface {
	// Create(newsfeed *dto.NewsfeedDTO) error
	// Delete(newsfeedID string) error
	// Get(newsfeedID string) (*dto.NewsfeedDTO, error)
	// All(userID string) ([]*dto.NewsfeedDTO, error)
	// AddArticle(newsfeedID, articleID string) error
	InsertArticle(article *dto.ArticleDTO) error
	// RemoveArticle(newsfeedID, articleID string) error
}
