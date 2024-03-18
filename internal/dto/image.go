package dto

import (
	"encoding/json"

	"newser.app/internal/domain/entity"
)

type ImageDTO struct {
	ID    entity.ID `json:"id,omitempty" db:"id"`
	URL   string    `json:"url,omitempty" db:"url"`
	Title string    `json:"title,omitempty" db:"title"`
}

func (i ImageDTO) JSON() []byte {
	j, _ := json.MarshalIndent(i, "", "  ")
	return j
}

func (i ImageDTO) String() string {
	return string(i.JSON())
}

func (i ImageDTO) FromDomain(image *entity.Image) ImageDTO {
	return ImageDTO{
		ID:    image.ID,
		URL:   image.URL.String(),
		Title: image.Title,
	}
}
