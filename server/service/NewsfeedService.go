package service

import (
	"newser.app/infra/repository"
	"newser.app/model"
)

type NewsfeedService struct {
	articleRepo repository.ArticleRepository
	imageRepo   repository.ImageRepository
	personRepo  repository.PersonRepository
}

func NewNewsfeedService(
	ar repository.ArticleRepository,
	ir repository.ImageRepository,
	pr repository.PersonRepository,
) NewsfeedService {
	return NewsfeedService{
		articleRepo: ar,
		imageRepo:   ir,
		personRepo:  pr,
	}
}

func (s *NewsfeedService) GetArticlesByNewsfeedId(nId int64) ([]*model.Article, error) {
	aa, err := s.articleRepo.ArticlesByNewsfeed(nId)
	if err != nil {
		return nil, err
	}
	return aa, nil
}

func (s *NewsfeedService) GetArticleById(aId int64) (*model.Article, error) {
	a, err := s.articleRepo.Get(aId)
	if err != nil {
		return nil, err
	}
	return a, nil
}
