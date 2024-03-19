package mapper

import (
	"fmt"
	"time"

	"github.com/mmcdole/gofeed"
	"newser.app/internal/domain/entity"
)

type GofeedMapper struct{}

func (m GofeedMapper) ToNewsfeed(gfFeed *gofeed.Feed) (*entity.Newsfeed, error) {
	nf, err := entity.NewNewsfeed(
		gfFeed.Title,
		gfFeed.Link,
		gfFeed.FeedLink,
		gfFeed.Description,
		gfFeed.Copyright,
		gfFeed.Language,
		gfFeed.FeedType,
		m.ToImage(gfFeed.Image),
	)

	if err != nil {
		return nil, err
	}

	for _, gfItem := range gfFeed.Items {
		article, err := m.ToArticle(gfItem)
		if err != nil {
			return nil, err
		}
		nf.Articles = append(nf.Articles, article)
	}
	return nf, nil
}

func (m GofeedMapper) ToArticle(gfItem *gofeed.Item) (*entity.Article, error) {
	var updateTime time.Time
	var publishedTime time.Time
	if gfItem.UpdatedParsed != nil {
		updateTime = *gfItem.UpdatedParsed
	}
	if gfItem.PublishedParsed != nil {
		publishedTime = *gfItem.PublishedParsed
	}
	article, err := entity.NewArticle(
		gfItem.Title,
		gfItem.Description,
		gfItem.Content,
		gfItem.Link,
		gfItem.Updated,
		gfItem.Published,
		gfItem.GUID,
		updateTime,
		publishedTime,
		m.ToPerson(gfItem.Author),
		m.ToImage(gfItem.Image),
		gfItem.Categories,
	)
	if err != nil {
		fmt.Println("error creating article: ", err)
		return nil, err
	}
	return article, nil
}

func (m GofeedMapper) ToImage(gfImage *gofeed.Image) *entity.Image {
	if gfImage == nil {
		return nil
	}
	return entity.NewImage(gfImage.URL, gfImage.Title)
}

func (m GofeedMapper) ToPerson(gfPerson *gofeed.Person) *entity.Person {
	if gfPerson == nil {
		return nil
	}
	ep, err := entity.NewPerson(gfPerson.Name, gfPerson.Email)
	if err != nil {
		return nil
	}
	return ep
}

func (m GofeedMapper) ToCategory(gfCategory string) *entity.Category {
	return entity.NewCategory(gfCategory)
}
