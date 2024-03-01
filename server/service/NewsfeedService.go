package service

import (
	"newser.app/infra/repository"
	"newser.app/model"
)

type NewsfeedService struct {
	articleRepo repository.ArticleRepository
}

func (s *NewsfeedService) GetArticlesByNewsfeedId(nId int64) ([]*model.Article, error) {
	stored, err := s.articleRepo.ArticlesByNewsfeed(nId)
	if err != nil {
		return nil, err
	}
	return stored, nil
}
