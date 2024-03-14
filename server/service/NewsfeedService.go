package service

import (
	"newser.app/infra/dto"
	"newser.app/infra/repository"
)

type NewsfeedService struct {
	articleRepo    repository.ArticleRepository
	imageRepo      repository.ImageRepository
	personRepo     repository.PersonRepository
	newsfeedRepo   repository.NewsfeedRepository
	collectionRepo repository.CollectionRepository
}

func NewNewsfeedService(
	ar repository.ArticleRepository,
	ir repository.ImageRepository,
	pr repository.PersonRepository,
	nr repository.NewsfeedRepository,
	cr repository.CollectionRepository,
) NewsfeedService {
	return NewsfeedService{
		articleRepo:    ar,
		imageRepo:      ir,
		personRepo:     pr,
		newsfeedRepo:   nr,
		collectionRepo: cr,
	}
}

func (s NewsfeedService) GetNewsfeed(id, userId int64) (*dto.NewsfeedDTO, error) {
	feed, err := s.newsfeedRepo.Get(id)
	if err != nil {
		return nil, err
	}

	feedDTO := dto.NewsfeedDTOFromDomain(feed)

	for _, a := range feedDTO.Articles {
		userReadCollection, err := s.collectionRepo.FindByTitle("read", userId)
		if err != nil {
			return nil, err
		}

		isRead, err := s.collectionRepo.IsArticleInCollection(a.ID, userReadCollection.Id)
		if err != nil {
			return nil, err
		}
		a.Read = isRead
	}

	return feedDTO, nil
}

func (s *NewsfeedService) GetArticlesByNewsfeedId(nId int64) ([]*dto.ArticleDTO, error) {
	aa, err := s.articleRepo.ArticlesByNewsfeed(nId)
	if err != nil {
		return nil, err
	}

	articleDTOs := []*dto.ArticleDTO{}
	for _, a := range aa {
		articleDTOs = append(articleDTOs, dto.ArticleDTOFromDomain(a))
	}
	return articleDTOs, nil
}

func (s *NewsfeedService) GetArticleById(aId int64) (*dto.ArticleDTO, error) {
	a, err := s.articleRepo.Get(aId)
	if err != nil {
		return nil, err
	}
	return dto.ArticleDTOFromDomain(a), nil
}
