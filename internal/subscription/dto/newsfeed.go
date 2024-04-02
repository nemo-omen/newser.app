package dto

import (
	"newser.app/internal/domain/entity"
)

type NewsfeedDTO struct {
	ID          string        `json:"id,omitempty" db:"id"`
	Title       string        `json:"title,omitempty" db:"title"`
	Description string        `json:"description,omitempty" db:"description"`
	FeedURL     string        `json:"feedURL,omitempty" db:"feed_url"`
	SiteURL     string        `json:"siteURL,omitempty" db:"site_url"`
	Language    string        `json:"language,omitempty" db:"language"`
	ImageTitle  string        `json:"image,omitempty" db:"feed_image_title"`
	ImageURL    string        `json:"imageURL,omitempty" db:"feed_image_url"`
	Copyright   string        `json:"copyRight,omitempty" db:"copyright"`
	Articles    []*ArticleDTO `json:"articles,omitempty" db:"articles"`
	FeedType    string        `json:"feedType,omitempty" db:"feed_type"`
	Slug        string        `json:"slug,omitempty" db:"slug"`
}

func (nf NewsfeedDTO) FromDomain(newsfeed entity.Newsfeed) NewsfeedDTO {
	articles := []*ArticleDTO{}
	for _, a := range newsfeed.Articles {
		articles = append(articles, ArticleDTO{}.FromDomain(a))
	}
	return NewsfeedDTO{
		ID:          newsfeed.ID.String(),
		Title:       newsfeed.Title,
		SiteURL:     newsfeed.SiteURL.String(),
		FeedURL:     newsfeed.FeedURL.String(),
		Description: newsfeed.Description,
		Copyright:   newsfeed.Copyright,
		Language:    newsfeed.Language,
		ImageTitle:  newsfeed.Image.Title,
		ImageURL:    newsfeed.Image.URL.String(),
		Articles:    articles,
		FeedType:    newsfeed.FeedType,
		Slug:        newsfeed.Slug.String(),
	}
}
