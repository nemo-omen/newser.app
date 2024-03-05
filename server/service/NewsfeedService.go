package service

import (
	"newser.app/infra/repository"
	"newser.app/model"
)

type NewsfeedService struct {
	articleRepo  repository.ArticleRepository
	imageRepo    repository.ImageRepository
	personRepo   repository.PersonRepository
	newsfeedRepo repository.NewsfeedRepository
}

func NewNewsfeedService(
	ar repository.ArticleRepository,
	ir repository.ImageRepository,
	pr repository.PersonRepository,
	nr repository.NewsfeedRepository,
) NewsfeedService {
	return NewsfeedService{
		articleRepo:  ar,
		imageRepo:    ir,
		personRepo:   pr,
		newsfeedRepo: nr,
	}
}

func (s NewsfeedService) GetNewsfeed(id int64) (*model.Newsfeed, error) {
	feed, err := s.newsfeedRepo.Get(id)
	if err != nil {
		return nil, err
	}
	articles, err := s.GetArticlesByNewsfeedId(feed.ID)
	if err != nil {
		return nil, err
	}
	feed.Articles = articles
	return feed, nil
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
