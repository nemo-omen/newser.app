package dto

import (
	"encoding/json"

	"newser.app/domain/entity"
)

type ArticleDTO struct {
	ID          entity.ID     `json:"id,omitempty"`
	Title       string        `json:"title,omitempty"`
	Description string        `json:"description,omitempty"`
	Content     string        `json:"content,omitempty"`
	Link        string        `json:"link,omitempty"`
	Updated     string        `json:"updated,omitempty"`
	Published   string        `json:"published,omitempty"`
	Author      PersonDTO     `json:"author,omitempty"`
	GUID        string        `json:"guid,omitempty"`
	Image       ImageDTO      `json:"image,omitempty"`
	Categories  []CategoryDTO `json:"categories,omitempty"`
	Read        bool          `json:"read,omitempty"`
	Saved       bool          `json:"saved,omitempty"`
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
		ID:          article.ID,
		Title:       article.Title,
		Description: article.Description,
		Content:     article.Content,
		Link:        article.Link.String(),
		Updated:     article.Updated,
		Published:   article.Published,
		Author:      PersonDTO{}.FromDomain(article.Author),
		GUID:        article.GUID,
		Image:       ImageDTO{}.FromDomain(article.Image),
		Categories:  []CategoryDTO{},
		Read:        article.Read,
		Saved:       article.Saved,
	}
	for _, c := range article.Categories {
		art.Categories = append(art.Categories, CategoryDTO{}.FromDomain(c))
	}
	return art
}
