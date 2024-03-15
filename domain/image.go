package domain

import (
	"github.com/google/uuid"
)

type Image struct {
	ID    uuid.UUID
	URL   string
	Title string
}

type ImageDTO struct {
	ID    uuid.UUID `json:"id,omitempty"`
	URL   string    `json:"url,omitempty"`
	Title string    `json:"title,omitempty"`
}

type ImageDAO struct {
	ID    uuid.UUID `db:"id"`
	URL   string    `db:"url"`
	Title string    `db:"title"`
}

func NewImage(url, title string) Image {
	return Image{
		ID:    uuid.New(),
		URL:   url,
		Title: title,
	}
}
