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

type NewsfeedDAO struct {
	ID          int64  `db:"id"`
	Title       string `db:"title"`
	SiteUrl     string `db:"site_url"`
	FeedUrl     string `db:"feed_url"`
	Description string `db:"description"`
	Copyright   string `db:"copyright"`
	AuthorId    int64  `db:"author_id"`
	Language    string `db:"language"`
	FeedType    string `db:"feed_type"`
	Slug        string `db:"slug"`
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

func NewsfeedDAOFromDomain(nf *model.Newsfeed) *NewsfeedDAO {
	return &NewsfeedDAO{
		ID:          nf.ID,
		Title:       nf.Title,
		SiteUrl:     nf.SiteUrl,
		FeedUrl:     nf.FeedUrl,
		Description: nf.Description,
		Copyright:   nf.Copyright,
		AuthorId:    nf.Author.ID,
		Language:    nf.Language,
		FeedType:    nf.FeedType,
		Slug:        nf.Slug,
	}
}

func (dto *NewsfeedDTO) ToDomain() *model.Newsfeed {
	articles := make([]*model.Article, len(dto.Articles))
	for i, a := range dto.Articles {
		articles[i] = a.ToDomain()
	}
	return &model.Newsfeed{
		ID:          dto.ID,
		Title:       dto.Title,
		SiteUrl:     dto.SiteUrl,
		FeedUrl:     dto.FeedUrl,
		Description: dto.Description,
		Image:       dto.Image,
		Categories:  dto.Categories,
		Articles:    articles,
		Author:      dto.Author,
		Language:    dto.Language,
		FeedType:    dto.FeedType,
		Slug:        dto.Slug,
	}
}
