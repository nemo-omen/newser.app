package dto

import (
	"encoding/json"

	"newser.app/domain/entity"
)

type ImageDTO struct {
	ID    entity.ID `json:"id,omitempty"`
	URL   string    `json:"url,omitempty"`
	Title string    `json:"title,omitempty"`
}

func (i ImageDTO) JSON() []byte {
	j, _ := json.MarshalIndent(i, "", "  ")
	return j
}

func (i ImageDTO) String() string {
	return string(i.JSON())
}

func (i ImageDTO) FromDomain(image entity.Image) ImageDTO {
	return ImageDTO{
		ID:    image.ID,
		URL:   image.URL,
		Title: image.Title,
	}
}
