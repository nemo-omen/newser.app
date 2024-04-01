package dto

import (
	"time"

	"newser.app/internal/search/entity"
)

type SearchResultFeedDTO struct {
	Title           string                   `json:"title,omitempty"`
	Description     string                   `json:"description,omitempty"`
	Link            string                   `json:"link,omitempty"`
	FeedLink        string                   `json:"feedLink,omitempty"`
	Links           []string                 `json:"links,omitempty"`
	Updated         string                   `json:"updated,omitempty"`
	UpdatedParsed   *time.Time               `json:"updatedParsed,omitempty"`
	Published       string                   `json:"published,omitempty"`
	PublishedParsed *time.Time               `json:"publishedParsed,omitempty"`
	Authors         []*SearchResultPersonDTO `json:"authors,omitempty"`
	Language        string                   `json:"language,omitempty"`
	Image           *SearchResultImageDTO    `json:"image,omitempty"`
	Copyright       string                   `json:"copyright,omitempty"`
	Generator       string                   `json:"generator,omitempty"`
	Categories      []string                 `json:"categories,omitempty"`
	Items           []*SearchResultItemDTO   `json:"items"`
	FeedType        string                   `json:"feedType"`
	FeedVersion     string                   `json:"feedVersion"`
}

func (f *SearchResultFeedDTO) ToDomain() *entity.SearchResultFeed {
	if f == nil {
		return nil
	}

	return &entity.SearchResultFeed{
		Title:           f.Title,
		Description:     f.Description,
		Link:            f.Link,
		FeedLink:        f.FeedLink,
		Links:           f.Links,
		Updated:         f.Updated,
		UpdatedParsed:   f.UpdatedParsed,
		Published:       f.Published,
		PublishedParsed: f.PublishedParsed,
		Authors:         f.Authors.ToDomain(),
		Language:        f.Language,
		Image:           f.Image.ToDomain(),
		Categories:      f.Categories,
		Items:           f.Items.ToDomain(),
		FeedType:        f.FeedType,
		FeedVersion:     f.FeedVersion,
	}
}
