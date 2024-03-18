package dto

import "newser.app/internal/domain/entity"

type NewsfeedDTO struct {
	ID          string       `json:"id,omitempty" db:"id"`
	Title       string       `json:"title,omitempty" db:"title"`
	Description string       `json:"description,omitempty" db:"description"`
	FeedLink    string       `json:"feedLink,omitempty" db:"feedLink"`
	SiteLink    string       `json:"siteLink,omitempty" db:"siteLink"`
	Author      PersonDTO    `json:"author,omitempty" db:"author"`
	Language    string       `json:"language,omitempty" db:"language"`
	Image       ImageDTO     `json:"image,omitempty" db:"image"`
	Copyright   string       `json:"copyRight,omitempty" db:"copyRight"`
	Articles    []ArticleDTO `json:"articles,omitempty" db:"articles"`
}

func (nf NewsfeedDTO) FromDomain(newsfeed entity.Newsfeed) NewsfeedDTO {
	articles := []ArticleDTO{}
	for _, a := range newsfeed.Articles {
		articles = append(articles, ArticleDTO{}.FromDomain(a))
	}
	return NewsfeedDTO{
		ID:          newsfeed.ID.String(),
		Title:       newsfeed.Title,
		Description: newsfeed.Description,
		FeedLink:    newsfeed.FeedLink.String(),
		SiteLink:    newsfeed.SiteLink.String(),
		Author:      PersonDTO{}.FromDomain(newsfeed.Author),
		Language:    newsfeed.Language,
		Image:       ImageDTO{}.FromDomain(newsfeed.Image),
		Copyright:   newsfeed.Copyright,
		Articles:    articles,
	}
}
