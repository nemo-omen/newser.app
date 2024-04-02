package entity

import (
	"encoding/json"

	"newser.app/internal/subscription/dto"
)

type Image struct {
	ID    int64
	URL   string
	Title string
}

func NewImage(data dto.ImageDTO) *Image {
	return &Image{
		ID:    data.ID,
		URL:   data.URL,
		Title: data.Title,
	}
}

func (i Image) JSON() []byte {
	json, _ := json.MarshalIndent(i, "", "  ")
	return json
}

func (i Image) String() string {
	return string(i.JSON())
}
