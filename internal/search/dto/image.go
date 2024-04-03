package dto

import "newser.app/internal/search/entity"

type ImageDTO struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

func (s *ImageDTO) ToDomain() *entity.Image {
	if s == nil {
		return nil
	}

	return &entity.Image{
		URL:   s.URL,
		Title: s.Title,
	}
}

func ImageToDTO(s *entity.Image) *ImageDTO {
	if s == nil {
		return nil
	}

	return &ImageDTO{
		URL:   s.URL,
		Title: s.Title,
	}
}
