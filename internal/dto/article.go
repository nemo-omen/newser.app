package dto

import (
	"encoding/json"
	"time"

	"newser.app/internal/domain/entity"
)

type ArticleDTO struct {
	ID              entity.ID     `json:"id,omitempty" db:"id"`
	Title           string        `json:"title,omitempty" db:"title"`
	Description     string        `json:"description,omitempty" db:"description"`
	Content         string        `json:"content,omitempty" db:"content"`
	Link            string        `json:"link,omitempty" db:"link"`
	Updated         string        `json:"updated,omitempty" db:"updated"`
	UpdatedParsed   time.Time     `json:"updated_parsed,omitempty" db:"updated_parsed"`
	Published       string        `json:"published,omitempty" db:"published"`
	PublishedParsed time.Time     `json:"published_parsed,omitempty" db:"published_parsed"`
	Author          PersonDTO     `json:"author,omitempty" db:"author"`
	GUID            string        `json:"guid,omitempty" db:"guid"`
	Image           ImageDTO      `json:"image,omitempty" db:"image"`
	Categories      []CategoryDTO `json:"categories,omitempty" db:"categories"`
	Read            bool          `json:"read,omitempty" db:"read"`
	Saved           bool          `json:"saved,omitempty" db:"saved"`
	SiteURL         string        `json:"site_url,omitempty" db:"site_url"`
	FeedTitle       string        `json:"feed_title,omitempty" db:"feed_title"`
	FeedImageURL    string        `json:"feed_image_url,omitempty" db:"feed_image_url"`
	FeedImageTitle  string        `json:"feed_image_title,omitempty" db:"feed_image_title"`
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
