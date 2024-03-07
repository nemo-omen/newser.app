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

func (s *CollectionService) AddArticleToCollectionByName(collectionName string, articleId int64) error {
	collection, err := s.collectionRepo.FindByTitle(collectionName)
	if err != nil {
		return err
	}

	err = s.collectionRepo.InsertCollectionItem(articleId, collection.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionService) RemoveArticleFromCollectionByName(collectionName string, articleId int64) error {
	collection, err := s.collectionRepo.FindByTitle(collectionName)
	if err != nil {
		return err
	}

	err = s.collectionRepo.DeleteCollectionItem(articleId, collection.Id)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionService) AddArticleToRead(articleId int64) error {
	article, err := s.articleRepo.Get(articleId)
	if err != nil {
		return err
	}

	err = s.AddArticleToCollectionByName("read", articleId)
	if err != nil {
		return err
	}
	err = s.RemoveArticleFromCollectionByName("unread", articleId)
	if err != nil {
		return err
	}
	article.Read = true
	_, err = s.articleRepo.Update(article)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionService) RemoveArticleFromRead(articleId int64) error {
	article, err := s.articleRepo.Get(articleId)
	if err != nil {
		return err
	}
	err = s.RemoveArticleFromCollectionByName("read", articleId)
	if err != nil {
		return err
	}
	err = s.AddArticleToCollectionByName("unread", articleId)
	if err != nil {
		return err
	}

	article.Read = false
	_, err = s.articleRepo.Update(article)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionService) AddArticleToSaved(articleId int64) error {
	err := s.AddArticleToCollectionByName("saved", articleId)
	if err != nil {
		return err
	}
	return nil
}

func (s *CollectionService) RemoveArticleFromSaved(articleId int64) error {
	err := s.RemoveArticleFromCollectionByName("saved", articleId)
	if err != nil {
		return err
	}
	return nil
}
