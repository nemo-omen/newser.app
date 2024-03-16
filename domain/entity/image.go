package entity

import (
	"encoding/json"
)

type Image struct {
	ID    ID     `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

func NewImage(url, title string) *Image {
	return &Image{
		ID:    NewID(),
		URL:   url,
		Title: title,
	}
}

func (i Image) JSON() []byte {
	json, _ := json.MarshalIndent(i, "", "  ")
	return json
}

func (i Image) String() string {
	return string(i.JSON())
}
