package dto

import (
	"encoding/json"
	"time"

	"newser.app/internal/search/entity"
)

type FeedDTO struct {
	Title           string       `json:"title,omitempty"`
	Description     string       `json:"description,omitempty"`
	Link            string       `json:"link,omitempty"`
	FeedLink        string       `json:"feedLink,omitempty"`
	Links           []string     `json:"links,omitempty"`
	Updated         string       `json:"updated,omitempty"`
	UpdatedParsed   *time.Time   `json:"updatedParsed,omitempty"`
	Published       string       `json:"published,omitempty"`
	PublishedParsed *time.Time   `json:"publishedParsed,omitempty"`
	Authors         []*PersonDTO `json:"authors,omitempty"`
	Language        string       `json:"language,omitempty"`
	Image           *ImageDTO    `json:"image,omitempty"`
	Copyright       string       `json:"copyright,omitempty"`
	Categories      []string     `json:"categories,omitempty"`
	Items           []*ItemDTO   `json:"items"`
	FeedType        string       `json:"feedType"`
}

func (f *FeedDTO) ToDomain() *entity.Feed {
	if f == nil {
		return nil
	}

	authors := []*entity.Person{}
	for _, author := range f.Authors {
		authors = append(authors, author.ToDomain())
	}

	items := []*entity.Item{}
	for _, item := range f.Items {
		items = append(items, item.ToDomain())
	}

	return &entity.Feed{
		Title:           f.Title,
		Description:     f.Description,
		Link:            f.Link,
		FeedLink:        f.FeedLink,
		Links:           f.Links,
		Updated:         f.Updated,
		UpdatedParsed:   f.UpdatedParsed,
		Published:       f.Published,
		PublishedParsed: f.PublishedParsed,
		Authors:         authors,
		Language:        f.Language,
		Image:           f.Image.ToDomain(),
		Categories:      f.Categories,
		Items:           items,
		FeedType:        f.FeedType,
	}
}

func FeedToDTO(feed *entity.Feed) *FeedDTO {
	if feed == nil {
		return nil
	}

	authors := []*PersonDTO{}
	for _, author := range feed.Authors {
		authors = append(authors, PersonToDTO(author))
	}

	items := []*ItemDTO{}
	for _, item := range feed.Items {
		items = append(items, ItemToDTO(item))
	}

	return &FeedDTO{
		Title:           feed.Title,
		Description:     feed.Description,
		Link:            feed.Link,
		FeedLink:        feed.FeedLink,
		Links:           feed.Links,
		Updated:         feed.Updated,
		UpdatedParsed:   feed.UpdatedParsed,
		Published:       feed.Published,
		PublishedParsed: feed.PublishedParsed,
		Authors:         authors,
		Language:        feed.Language,
		Image:           ImageToDTO(feed.Image),
		Categories:      feed.Categories,
		Items:           items,
		FeedType:        feed.FeedType,
	}
}

func (d *FeedDTO) JSON() []byte {
	j, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return []byte{}
	}
	return j
}
