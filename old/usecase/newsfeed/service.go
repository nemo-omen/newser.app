package newsfeed

import (
	"github.com/mmcdole/gofeed"
	"newser.app/internal/dto"
	"newser.app/internal/infra/mapper"
	"newser.app/internal/infra/repository"
)

type NewsfeedService struct {
	newsfeedRepo repository.NewsfeedRepository
}

func NewNewsfeedService(newsfeedRepo repository.NewsfeedRepository) NewsfeedService {
	return NewsfeedService{
		newsfeedRepo: newsfeedRepo,
	}
}

func (s *NewsfeedService) SaveArticle(item *gofeed.Item) error {
	mapper := mapper.GofeedMapper{}
	articleEntity, err := mapper.ToArticle(item)
	if err != nil {
		return err
	}
	articleDTO := dto.ArticleDTO{}.FromDomain(articleEntity)
	return s.newsfeedRepo.InsertArticle(articleDTO)
}
