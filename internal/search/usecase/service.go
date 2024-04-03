package usecase

import (
	"newser.app/internal/search"
	"newser.app/internal/search/dto"
)

// SearchService provides search operations.
type SearchService struct {
	repo search.SearchRepository
}

// NewSearchService creates a new SearchService.
func NewSearchService(repo search.SearchRepository) *SearchService {
	return &SearchService{
		repo: repo,
	}
}

func (s *SearchService) FindFeedUrls(url string) ([]string, error) {
	return s.repo.FindFeedLinks(url)
}

func (s *SearchService) GetFeed(url string) (*dto.FeedDTO, error) {
	feed, err := s.repo.GetFeed(url)
	if err != nil {
		return nil, err
	}
	return dto.FeedToDTO(feed), nil
}

func (s *SearchService) GetFeeds(urls []string) ([]*dto.FeedDTO, error) {
	feeds, err := s.repo.GetFeeds(urls)
	if err != nil {
		return nil, err
	}
	var feedDTOs []*dto.FeedDTO
	for _, feed := range feeds {
		feedDTOs = append(feedDTOs, dto.FeedToDTO(feed))
	}
	return feedDTOs, nil
}
