package service

import (
	"newser.app/infra/repository"
	"newser.app/model"
)

type CollectionService struct {
	collectionRepo repository.CollectionRepository
	articleRepo    repository.ArticleRepository
}

func NewCollectionService(cr repository.CollectionRepository, ar repository.ArticleRepository) CollectionService {
	return CollectionService{
		collectionRepo: cr,
		articleRepo:    ar,
	}
}

func (s *CollectionService) GetArticlesByCollectionByName(cName string, userId int64) ([]*model.Article, error) {
	collectionArticles, err := s.collectionRepo.GetArticlesByCollectionName(cName, userId)
	if err != nil {
		return nil, err
	}
	return collectionArticles, nil
}

func (s *CollectionService) AddArticleToCollectionByName(collectionName string, articleId, userId int64) (*model.Article, error) {
	collection, err := s.collectionRepo.FindByTitle(collectionName, userId)
	if err != nil {
		return nil, err
	}

	article, err := s.articleRepo.Get(articleId)
	if err != nil {
		return nil, err
	}

	err = s.collectionRepo.InsertCollectionItem(articleId, collection.Id)
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (s *CollectionService) RemoveArticleFromCollectionByName(collectionName string, articleId, userId int64) error {
	collection, err := s.collectionRepo.FindByTitle(collectionName, userId)
	if err != nil {
		return err
	}

	err = s.collectionRepo.DeleteCollectionItem(articleId, collection.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionService) AddArticleToSaved(articleId, userId int64) (*model.Article, error) {
	article, err := s.AddArticleToCollectionByName("saved", articleId, userId)
	if err != nil {
		return nil, err
	}
	article.Saved = true
	return article, nil
}

func (s *CollectionService) RemoveArticleFromSaved(articleId, userId int64) error {
	article, err := s.articleRepo.Get(articleId)
	if err != nil {
		return err
	}
	err = s.RemoveArticleFromCollectionByName("saved", articleId, userId)
	if err != nil {
		return err
	}
	article.Saved = true
	return nil
}

func (s *CollectionService) AddArticleToRead(articleId, userId int64) error {
	err := s.collectionRepo.MarkArticleRead(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionService) RemoveArticleFromRead(articleId, userId int64) error {
	err := s.collectionRepo.MarkArticleUnread(articleId, userId)
	if err != nil {
		return err
	}
	return nil
}
