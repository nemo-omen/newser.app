package dto

import "newser.app/model"

type NewsfeedDTO struct {
	ID          int64         `json:"id"`
	Title       string        `json:"title"`
	SiteUrl     string        `json:"site_url"`
	FeedUrl     string        `json:"feed_url"`
	Description string        `json:"description"`
	Image       *model.Image  `json:"image"`
	Copyright   string        `json:"copyright"`
	Articles    []*ArticleDTO `json:"articles"`
	Author      *model.Person `json:"author"`
	Language    string        `json:"language"`
	Categories  []string      `json:"categories"`
	FeedType    string        `json:"feed_type"`
	Slug        string        `json:"slug"`
}

func NewsfeedDTOFromDomain(nf *model.Newsfeed) *NewsfeedDTO {
	articles := make([]*ArticleDTO, len(nf.Articles))
	for i, a := range nf.Articles {
		articles[i] = ArticleDTOFromDomain(a)
	}
	return &NewsfeedDTO{
		ID:          nf.ID,
		Title:       nf.Title,
		SiteUrl:     nf.SiteUrl,
		FeedUrl:     nf.FeedUrl,
		Description: nf.Description,
		Image:       nf.Image,
		Copyright:   nf.Copyright,
		Articles:    articles,
		Author:      nf.Author,
		Language:    nf.Language,
		Categories:  nf.Categories,
		FeedType:    nf.FeedType,
		Slug:        nf.Slug,
	}
}
