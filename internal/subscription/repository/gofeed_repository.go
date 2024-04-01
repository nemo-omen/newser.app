package repository

import (
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
	"newser.app/internal/search/dto"
)

type GofeedRepository struct {
	client *http.Client
	parser *gofeed.Parser
}

func NewGoFeedRepository() *GofeedRepository {
	return &GofeedRepository{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		parser: gofeed.NewParser(),
	}
}

func (g *GofeedRepository) FindFeedUrls(url string) ([]string, error) {
	return nil, nil
}

func (g *GofeedRepository) GetFeed(url string) (*dto.SearchResultFeedDTO, error) {
	return nil, nil
}

func (g *GofeedRepository) GetFeeds(urls []string) ([]*dto.SearchResultFeedDTO, error) {
	return nil, nil
}
