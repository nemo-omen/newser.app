package entity

import (
	"encoding/json"

	"newser.app/domain/value"
)

type Newsfeed struct {
	ID          ID         `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	FeedLink    value.Link `json:"feedLink"`
	SiteLink    value.Link `json:"siteLink"`
	Author      *Person    `json:"author"`
	Language    string     `json:"language"`
	Image       *Image     `json:"image"`
	Copyright   string     `json:"copyRight"`
	Articles    []*Article `json:"articles"`
}

func NewNewsfeed(title, description, feedLink, siteLink, language, copyRight string, author *Person, image *Image) *Newsfeed {
	validFeedLink, err := value.NewLink(feedLink)
	if err != nil {
		return nil
	}
	validSiteLink, err := value.NewLink(siteLink)
	if err != nil {
		return nil
	}
	return &Newsfeed{
		ID:          NewID(),
		Title:       title,
		Description: description,
		FeedLink:    validFeedLink,
		SiteLink:    validSiteLink,
		Author:      author,
		Language:    language,
		Image:       image,
		Articles:    []*Article{},
	}
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
