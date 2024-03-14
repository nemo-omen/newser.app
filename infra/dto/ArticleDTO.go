package dto

import (
	"time"

	"newser.app/model"
)

// DTO -> Client
// DTO should have more information than is generally
// persisted to the database because it should carry
// all the information needed to render the view.
// This may include information retrieved from tables
// other than the primary table for the model the DTO
// represents/coincides with.

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

// DAO -> Database
// DAO should have the same fields as the database table
// it represents. It should be used to interact with the
// database and should not be exposed to the client.
// If we were working with an ORM, this would be the
// struct that would be mapped to the database table.
type ArticleDAO struct {
	ID              int64     `db:"id"`
	Title           string    `db:"title"`
	ArticleLink     string    `db:"article_link"`
	AuthorId        int64     `db:"author_id"`
	Description     string    `db:"description"`
	Content         string    `db:"content"`
	Published       string    `db:"published"`
	PublishedParsed time.Time `db:"published_parsed"`
	Updated         string    `db:"updated"`
	UpdatedParsed   time.Time `db:"updated_parsed"`
	GUID            string    `db:"guid"`
	Slug            string    `db:"slug"`
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

func ArticleDAOFromDomain(a *model.Article) *ArticleDAO {
	return &ArticleDAO{
		ID:              a.ID,
		Title:           a.Title,
		ArticleLink:     a.ArticleLink,
		AuthorId:        a.Person.ID,
		Description:     a.Description,
		Content:         a.Content,
		Published:       a.Published,
		PublishedParsed: a.PublishedParsed,
		Updated:         a.Updated,
		UpdatedParsed:   a.UpdatedParsed,
		GUID:            a.GUID,
		Slug:            a.Slug,
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

func (d ArticleDAO) ToDomain() *model.Article {
	return &model.Article{
		ID:              d.ID,
		Title:           d.Title,
		ArticleLink:     d.ArticleLink,
		Person:          model.Person{ID: d.AuthorId},
		Description:     d.Description,
		Content:         d.Content,
		Published:       d.Published,
		PublishedParsed: d.PublishedParsed,
		Updated:         d.Updated,
		UpdatedParsed:   d.UpdatedParsed,
		GUID:            d.GUID,
		Slug:            d.Slug,
	}
}
