package dto

import (
	"encoding/json"

	"newser.app/internal/domain/entity"
)

type ImageDTO struct {
	ID    int64  `json:"id,omitempty" db:"id"`
	URL   string `json:"url,omitempty" db:"url"`
	Title string `json:"title,omitempty" db:"title"`
}

func (i ImageDTO) JSON() []byte {
	j, _ := json.MarshalIndent(i, "", "  ")
	return j
}

func (i ImageDTO) String() string {
	return string(i.JSON())
}

func (i ImageDTO) FromDomain(image *entity.Image) ImageDTO {
	if image == nil {
		return ImageDTO{}
	}
	return ImageDTO{
		ID:    image.ID,
		URL:   image.URL.String(),
		Title: image.Title,
	}
}

func (i ImageDTO) ToDomain() *entity.Image {
	return &entity.Image{
		ID:    i.ID,
		URL:   entity.NewURL(i.URL),
		Title: i.Title,
	}
}
