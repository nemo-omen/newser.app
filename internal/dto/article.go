package dto

import (
	"encoding/json"
	"time"

	"newser.app/internal/domain/entity"
)

type ArticleDTO struct {
	ID              entity.ID     `json:"id,omitempty"`
	Title           string        `json:"title,omitempty"`
	Description     string        `json:"description,omitempty"`
	Content         string        `json:"content,omitempty"`
	Link            string        `json:"link,omitempty"`
	Updated         string        `json:"updated,omitempty"`
	UpdatedParsed   time.Time     `json:"updated_parsed,omitempty"`
	Published       string        `json:"published,omitempty"`
	PublishedParsed time.Time     `json:"published_parsed,omitempty"`
	Author          PersonDTO     `json:"author,omitempty"`
	GUID            string        `json:"guid,omitempty"`
	Image           ImageDTO      `json:"image,omitempty"`
	Categories      []CategoryDTO `json:"categories,omitempty"`
	Read            bool          `json:"read,omitempty"`
	Saved           bool          `json:"saved,omitempty"`
	SiteURL         string        `json:"site_url,omitempty"`
	FeedTitle       string        `json:"feed_title,omitempty"`
	FeedImageURL    string        `json:"feed_image_url,omitempty"`
	FeedImageTitle  string        `json:"feed_image_title,omitempty"`
}

func (a ArticleDTO) JSON() []byte {
	j, _ := json.MarshalIndent(a, "", "  ")
	return j
}

func (a ArticleDTO) String() string {
	return string(a.JSON())
}

func (a ArticleDTO) FromDomain(article *entity.Article) ArticleDTO {
	art := ArticleDTO{
		ID:              article.ID,
		Title:           article.Title,
		Description:     article.Description,
		Content:         article.Content,
		Link:            article.Link.String(),
		Updated:         article.Updated,
		UpdatedParsed:   article.UpdatedParsed,
		Published:       article.Published,
		PublishedParsed: article.PublishedParsed,
		Author:          PersonDTO{}.FromDomain(article.Author),
		GUID:            article.GUID,
		Image:           ImageDTO{}.FromDomain(article.Image),
		Categories:      []CategoryDTO{},
		Read:            article.Read,
		Saved:           article.Saved,
	}
	for _, c := range article.Categories {
		art.Categories = append(art.Categories, CategoryDTO{}.FromDomain(c))
	}
	return art
}
