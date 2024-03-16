package dto

import "newser.app/domain/entity"

type NewsfeedDTO struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	FeedLink    string       `json:"feedLink"`
	SiteLink    string       `json:"siteLink"`
	Author      PersonDTO    `json:"author"`
	Language    string       `json:"language"`
	Image       ImageDTO     `json:"image"`
	Copyright   string       `json:"copyRight"`
	Articles    []ArticleDTO `json:"articles"`
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
