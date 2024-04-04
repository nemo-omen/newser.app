package dto

import (
	"encoding/json"

	"newser.app/internal/subscription/entity"
)

type ImageDTO struct {
	ID    int64  `json:"id,omitempty"`
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
}

func (i ImageDTO) JSON() []byte {
	j, _ := json.MarshalIndent(i, "", "  ")
	return j
}

func (i ImageDTO) String() string {
	return string(i.JSON())
}

func (i ImageDTO) ToDomain() entity.Image {
	return entity.Image{
		ID:    i.ID,
		URL:   i.URL,
		Title: i.Title,
	}
}
