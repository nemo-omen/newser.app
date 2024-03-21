package collection

import (
	"newser.app/internal/dto"
	"newser.app/internal/infra/repository"
)

type CollectionService struct {
	collectionRepo repository.CollectionRepository
}

func NewCollectionService(collectionRepo repository.CollectionRepository) CollectionService {
	return CollectionService{
		collectionRepo: collectionRepo,
	}
}

func (s *CollectionService) GetArticlesBySlug(slug, userId string) ([]*dto.ArticleDTO, error) {
	collection, err := s.collectionRepo.GetBySlug(slug, userId)
	if err != nil {
		return nil, err
	}
	return s.collectionRepo.GetCollectionArticles(collection.ID)
}
