package dto

import "newser.app/internal/subscription/entity"

type SearchResultImageDTO struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

func (s *SearchResultImageDTO) ToDomain() *entity.SearchResultImage {
	if s == nil {
		return nil
	}

	return &entity.SearchResultImage{
		URL:   s.URL,
		Title: s.Title,
	}
}
