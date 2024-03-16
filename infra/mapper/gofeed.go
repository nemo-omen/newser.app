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
	return &entity.Article{
		Item: entity.Item{
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
			Categories:  gfItem.Categories,
		},
		Read:  false,
		Saved: false,
	}
}

func (m GofeedMapper) ToImage(gfImage *gofeed.Image) *entity.Image {
	return &entity.Image{
		ID:    entity.NewID(),
		URL:   gfImage.URL,
		Title: gfImage.Title,
	}
}

func (m GofeedMapper) ToPerson(gfPerson *gofeed.Person) *entity.Person {
	name, err := value.NewName(gfPerson.Name)
	if err != nil {
		name, _ = value.NewName("Unknown")
	}

	email, err := value.NewEmail(gfPerson.Email)
	if err != nil {
		email, _ = value.NewEmail("unknown@unknown.com")
	}

	return &entity.Person{
		ID:    entity.NewID(),
		Name:  name,
		Email: email,
	}
}
