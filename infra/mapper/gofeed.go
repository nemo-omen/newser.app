package mapper

import (
	"github.com/mmcdole/gofeed"
	"newser.app/domain/entity"
	"newser.app/domain/value"
)

type GofeedMapper struct{}

func (m GofeedMapper) ToNewsfeed(gfFeed *gofeed.Feed) *entity.Newsfeed {
	nf := &entity.Newsfeed{
		ID:          entity.NewID(),
		Title:       gfFeed.Title,
		Description: gfFeed.Description,
		FeedLink:    value.Link(gfFeed.FeedLink),
		SiteLink:    value.Link(gfFeed.Link),
		Author:      m.ToPerson(gfFeed.Author),
		Language:    gfFeed.Language,
		Image:       m.ToImage(gfFeed.Image),
		Copyright:   gfFeed.Copyright,
		Articles:    []*entity.Article{},
	}

	for _, gfItem := range gfFeed.Items {
		nf.Articles = append(nf.Articles, m.ToArticle(gfItem))
	}
	return nf
}

func (m GofeedMapper) ToArticle(gfItem *gofeed.Item) *entity.Article {
	categories := []*entity.Category{}
	for _, c := range gfItem.Categories {
		categories = append(categories, m.ToCategory(c))
	}
	return &entity.Article{
		Item: &entity.Item{
			ID:          entity.NewID(),
			Title:       gfItem.Title,
			Description: gfItem.Description,
			Content:     gfItem.Content,
			Link:        value.Link(gfItem.Link),
			Updated:     gfItem.Updated,
			Published:   gfItem.Published,
			Author:      m.ToPerson(gfItem.Author),
			GUID:        gfItem.GUID,
			Image:       m.ToImage(gfItem.Image),
			Categories:  categories,
		},
		Read:  false,
		Saved: false,
	}
}

func (m GofeedMapper) ToImage(gfImage *gofeed.Image) *entity.Image {
	validLink, err := value.NewLink(gfImage.URL)
	if err != nil {
		return nil
	}
	return &entity.Image{
		ID:    entity.NewID(),
		URL:   validLink,
		Title: gfImage.Title,
	}
}

func (m GofeedMapper) ToPerson(gfPerson *gofeed.Person) *entity.Person {
	name, err := value.NewName(gfPerson.Name)
	if err != nil {
		return nil
	}

	email, err := value.NewEmail(gfPerson.Email)
	if err != nil {
		email, _ = value.NewEmail("unknown@unknown.unknown")
	}

	return &entity.Person{
		ID:    entity.NewID(),
		Name:  name,
		Email: email,
	}
}

func (m GofeedMapper) ToCategory(gfCategory string) *entity.Category {
	return entity.NewCategory(gfCategory)
}
