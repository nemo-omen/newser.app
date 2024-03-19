package entity

import (
	"encoding/json"

	"newser.app/internal/domain/value"
	"newser.app/shared"
)

type Newsfeed struct {
	ID          ID         `json:"id"`
	Title       string     `json:"title"`
	SiteURL     value.Link `json:"siteLink"`
	FeedURL     value.Link `json:"feedLink"`
	Description string     `json:"description"`
	Copyright   string     `json:"copyRight"`
	Language    string     `json:"language"`
	Image       *Image     `json:"image"`
	Articles    []*Article `json:"articles"`
	FeedType    string     `json:"feedType"`
	Slug        value.Slug `json:"slug"`
}

func NewNewsfeed(
	title,
	siteLink,
	feedLink,
	description,
	copyright,
	language,
	feedType string,
	image *Image,
) (*Newsfeed, error) {
	validFeedLink, err := value.NewLink(feedLink)
	if err != nil {
		validFeedLink = ""
	}
	validSiteLink, err := value.NewLink(siteLink)
	if err != nil {
		validSiteLink = ""
	}
	slug, err := value.NewSlug(title)
	if err != nil {
		valErr, ok := err.(shared.AppError)
		if ok {
			valErr.Print()
		}
		return nil, err
	}
	return &Newsfeed{
		ID:          NewID(),
		Title:       title,
		SiteURL:     validSiteLink,
		FeedURL:     validFeedLink,
		Description: description,
		Copyright:   copyright,
		Language:    language,
		Image:       image,
		Articles:    []*Article{},
		FeedType:    feedType,
		Slug:        slug,
	}, nil
}

func (nf *Newsfeed) AddArticle(article *Article) {
	nf.Articles = append(nf.Articles, article)
}

func (nf *Newsfeed) AddArticles(articles []*Article) {
	nf.Articles = append(nf.Articles, articles...)
}

func (nf Newsfeed) JSON() []byte {
	j, _ := json.MarshalIndent(nf, "", "  ")
	return j
}

func (nf Newsfeed) String() string {
	return string(nf.JSON())
}
