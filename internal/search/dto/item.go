package dto

import (
	"time"

	"newser.app/internal/search/entity"
)

// ItemDTO is the object used to present fetched feeds to the
// client.
type ItemDTO struct {
	Title           string       `json:"title,omitempty"`
	Description     string       `json:"description,omitempty"`
	Content         string       `json:"content,omitempty"`
	Link            string       `json:"link,omitempty"`
	Links           []string     `json:"links,omitempty"`
	Updated         string       `json:"updated,omitempty"`
	UpdatedParsed   *time.Time   `json:"updatedParsed,omitempty"`
	Published       string       `json:"published,omitempty"`
	PublishedParsed *time.Time   `json:"publishedParsed,omitempty"`
	Authors         []*PersonDTO `json:"authors,omitempty"`
	GUID            string       `json:"guid,omitempty"`
	Image           *ImageDTO    `json:"image,omitempty"`
	Categories      []string     `json:"categories,omitempty"`
}

func (i *ItemDTO) ToDomain() *entity.Item {
	if i == nil {
		return nil
	}

	domainItem := &entity.Item{
		Title:           i.Title,
		Description:     i.Description,
		Content:         i.Content,
		Link:            i.Link,
		Links:           i.Links,
		Updated:         i.Updated,
		UpdatedParsed:   i.UpdatedParsed,
		Published:       i.Published,
		PublishedParsed: i.PublishedParsed,
		Authors:         make([]*entity.Person, 0, len(i.Authors)),
		GUID:            i.GUID,
		Image:           i.Image.ToDomain(),
		Categories:      i.Categories,
	}
	for _, author := range i.Authors {
		domainItem.Authors = append(domainItem.Authors, author.ToDomain())
	}
	return domainItem
}

// ItemToDTO converts an entity.Item to an ItemDTO.
func ItemToDTO(i *entity.Item) *ItemDTO {
	if i == nil {
		return nil
	}

	authors := []*PersonDTO{}
	for _, author := range i.Authors {
		authors = append(authors, PersonToDTO(author))
	}

	return &ItemDTO{
		Title:           i.Title,
		Description:     i.Description,
		Content:         i.Content,
		Link:            i.Link,
		Links:           i.Links,
		Updated:         i.Updated,
		UpdatedParsed:   i.UpdatedParsed,
		Published:       i.Published,
		PublishedParsed: i.PublishedParsed,
		Authors:         authors,
		GUID:            i.GUID,
		Image:           ImageToDTO(i.Image),
		Categories:      i.Categories,
	}
}
