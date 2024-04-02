package dto

import (
	"time"

	"newser.app/internal/subscription/entity"
)

// SearchResultItemDTO is the object used to present fetched feeds to the
// client.
type SearchResultItemDTO struct {
	Title           string                   `json:"title,omitempty"`
	Description     string                   `json:"description,omitempty"`
	Content         string                   `json:"content,omitempty"`
	Link            string                   `json:"link,omitempty"`
	Links           []string                 `json:"links,omitempty"`
	Updated         string                   `json:"updated,omitempty"`
	UpdatedParsed   *time.Time               `json:"updatedParsed,omitempty"`
	Published       string                   `json:"published,omitempty"`
	PublishedParsed *time.Time               `json:"publishedParsed,omitempty"`
	Authors         []*SearchResultPersonDTO `json:"authors,omitempty"`
	GUID            string                   `json:"guid,omitempty"`
	Image           *SearchResultImageDTO    `json:"image,omitempty"`
	Categories      []string                 `json:"categories,omitempty"`
}

func (i *SearchResultItemDTO) ToDomain() *entity.SearchResultItem {
	if i == nil {
		return nil
	}

	domainItem := &entity.SearchResultItem{
		Title:           i.Title,
		Description:     i.Description,
		Content:         i.Content,
		Link:            i.Link,
		Links:           i.Links,
		Updated:         i.Updated,
		UpdatedParsed:   i.UpdatedParsed,
		Published:       i.Published,
		PublishedParsed: i.PublishedParsed,
		Authors:         make([]*entity.SearchResultPerson, 0, len(i.Authors)),
		GUID:            i.GUID,
		Image:           i.Image.ToDomain(),
		Categories:      i.Categories,
	}
	for _, author := range i.Authors {
		domainItem.Authors = append(domainItem.Authors, author.ToDomain())
	}
	return domainItem
}
