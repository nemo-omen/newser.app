package dto

import (
	"time"

	"newser.app/model"
)

type ArticleDTO struct {
	ID              int64         `json:"id"`
	Title           string        `json:"title"`
	ArticleLink     string        `json:"link"`
	Author          *model.Person `json:"author"`
	Description     string        `json:"description"`
	Content         string        `json:"content"`
	Published       string        `json:"published"`
	PublishedParsed time.Time     `json:"published_parsed"`
	Updated         string        `json:"updated"`
	UpdatedParsed   time.Time     `json:"updated_parsed"`
	GUID            string        `json:"guid"`
	Slug            string        `json:"slug"`
	FeedId          int64         `json:"feed_id"`
	Read            bool          `json:"read"`
	Saved           bool          `json:"saved"`
	FeedUrl         string        `json:"feed_url"`
	FeedTitle       string        `json:"feed_title"`
	FeedSiteUrl     string        `json:"feed_site_url"`
	FeedSlug        string        `json:"feed_slug"`
	FeedImageUrl    string        `json:"feed_image_url"`
	FeedImageTitle  string        `json:"feed_image_title"`
}

func ArticleDTOFromDomain(a *model.Article) *ArticleDTO {
	return &ArticleDTO{
		ID:              a.ID,
		Title:           a.Title,
		ArticleLink:     a.ArticleLink,
		Description:     a.Description,
		Content:         a.Content,
		Author:          &a.Person,
		Published:       a.Published,
		PublishedParsed: a.PublishedParsed,
		Updated:         a.Updated,
		UpdatedParsed:   a.UpdatedParsed,
		GUID:            a.GUID,
		Slug:            a.Slug,
		FeedId:          a.FeedId,
	}
}

func (d ArticleDTO) ToDomain() *model.Article {
	return &model.Article{
		ID:          d.ID,
		Title:       d.Title,
		Person:      *d.Author,
		ArticleLink: d.ArticleLink,
		Description: d.Description,
		Content:     d.Content,
		Published:   d.Published,
		Updated:     d.Updated,
		GUID:        d.GUID,
		Slug:        d.Slug,
		FeedId:      d.FeedId,
	}
}
