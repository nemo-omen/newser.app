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

func (s *CollectionService) GetArticlesBySlug(slug, userID string) ([]*dto.ArticleDTO, error) {
	collection, err := s.collectionRepo.GetBySlug(slug, userID)
	if err != nil {
		return nil, err
	}
	return s.collectionRepo.GetCollectionArticles(collection.ID, userID)
}

func (s *CollectionService) AddArticleToCollection(collectionSlug, articleId, userId string) error {
	collection, err := s.collectionRepo.GetBySlug(collectionSlug, userId)
	if err != nil {
		return err
	}
	return s.collectionRepo.AddArticle(collection.ID, articleId)
}

func (s *CollectionService) RemoveArticleFromCollection(collectionSlug, articleId, userId string) error {
	collection, err := s.collectionRepo.GetBySlug(collectionSlug, userId)
	if err != nil {
		return err
	}
	return s.collectionRepo.RemoveArticle(collection.ID, articleId)
}

func (s *CollectionService) AddAndRemoveArticleFromCollection(addCollectionSlug, removeCollectionSlug, articleId, userId string) error {
	addCollection, err := s.collectionRepo.GetBySlug(addCollectionSlug, userId)
	if err != nil {
		return err
	}
	removeCollection, err := s.collectionRepo.GetBySlug(removeCollectionSlug, userId)
	if err != nil {
		return err
	}
	return s.collectionRepo.AddAndRemoveArticle(addCollection.ID, removeCollection.ID, articleId)
}
