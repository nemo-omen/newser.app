package dto

import (
	"encoding/json"
	"time"

	"newser.app/internal/subscription/entity"
)

type ArticleDTO struct {
	ID              int64     `json:"id,omitempty"`
	Title           string    `json:"title,omitempty"`
	Description     string    `json:"description,omitempty"`
	Content         string    `json:"content,omitempty"`
	Link            string    `json:"link,omitempty"`
	Updated         string    `json:"updated,omitempty"`
	UpdatedParsed   time.Time `json:"updated_parsed,omitempty"`
	Published       string    `json:"published,omitempty"`
	PublishedParsed time.Time `json:"published_parsed,omitempty"`
	Author          PersonDTO `json:"author,omitempty"`
	GUID            string    `json:"guid,omitempty"`
	Image           ImageDTO  `json:"image,omitempty"`
	Categories      []string  `json:"categories,omitempty"`
	Slug            string    `json:"slug,omitempty"`
	Read            bool      `json:"read" `
	Saved           bool      `json:"saved" `
	SiteURL         string    `json:"site_url,omitempty"`
	FeedTitle       string    `json:"feed_title,omitempty"`
	FeedID          string    `json:"feed_id,omitempty"`
	FeedImageURL    string    `json:"feed_image_url,omitempty"`
	FeedImageTitle  string    `json:"feed_image_title,omitempty"`
	FeedSlug        string    `json:"feed_slug,omitempty"`
}

func (a ArticleDTO) JSON() []byte {
	j, _ := json.MarshalIndent(a, "", "  ")
	return j
}

func (a ArticleDTO) String() string {
	return string(a.JSON())
}

func (a ArticleDTO) ToDomain() entity.Article {
	art := entity.Article{
		ID:              a.ID,
		Title:           a.Title,
		Description:     a.Description,
		Content:         a.Content,
		Link:            a.Link,
		Updated:         a.Updated,
		UpdatedParsed:   a.UpdatedParsed,
		Published:       a.Published,
		PublishedParsed: a.PublishedParsed,
		Author:          a.Author.ToDomain(),
		GUID:            a.GUID,
		Image:           a.Image.ToDomain(),
		Categories:      a.Categories,
		Slug:            a.Slug,
	}
	return art
}
