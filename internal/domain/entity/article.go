package entity

import (
	"encoding/json"
	"time"

	"newser.app/internal/domain/value"
	"newser.app/shared"
)

type Article struct {
	*Item
	Read  bool `json:"read"`
	Saved bool `json:"saved"`
}

type Item struct {
	ID              ID          `json:"id"`
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	Content         string      `json:"content"`
	Link            value.Link  `json:"link"`
	Updated         string      `json:"updated"`
	UpdatedParsed   time.Time   `json:"updated_parsed"`
	Published       string      `json:"published"`
	PublishedParsed time.Time   `json:"published_parsed"`
	Author          *Person     `json:"author"`
	GUID            string      `json:"guid"`
	Image           *Image      `json:"image"`
	Slug            value.Slug  `json:"slug"`
	Categories      []*Category `json:"categories"`
}

func NewArticle(
	title,
	description,
	content,
	link,
	updated,
	published,
	guid string,
	updatedParsed,
	publishedParsed time.Time,
	author *Person,
	image *Image,
	categories []string,
) (*Article, error) {
	validLink, err := value.NewLink(link)
	if err != nil {
		validLink = ""
	}
	slug, err := value.NewSlug(title)
	if err != nil {
		valErr, ok := err.(shared.AppError)
		if ok {
			valErr.Print()
		}
		return nil, err
	}

	a := &Article{
		Item: &Item{
			ID:              NewID(),
			Title:           title,
			Description:     description,
			Content:         content,
			Link:            validLink,
			Updated:         updated,
			UpdatedParsed:   updatedParsed,
			Published:       published,
			PublishedParsed: publishedParsed,
			Author:          author,
			GUID:            guid,
			Image:           image,
			Slug:            slug,
			Categories:      []*Category{},
		},
		Read:  false,
		Saved: false,
	}
	for _, c := range categories {
		a.Categories = append(a.Categories, NewCategory(c))
	}
	return a, nil
}

func (a Article) JSON() []byte {
	j, _ := json.MarshalIndent(a, "", "  ")
	return j
}

func (a Article) String() string {
	return string(a.JSON())
}

func (a *Article) SetRead() {
	a.Read = true
}

func (a *Article) SetSaved() {
	a.Saved = true
}

func (a *Article) SetUnread() {
	a.Read = false
}
