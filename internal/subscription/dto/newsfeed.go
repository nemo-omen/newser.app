package dto

"newser.app/internal/subscription/entity"

type NewsfeedDTO struct {
	ID          int64        `json:"id,omitempty"`
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	FeedURL     string       `json:"feedURL,omitempty"`
	SiteURL     string       `json:"siteURL,omitempty"`
	Language    string       `json:"language,omitempty"`
	Image       ImageDTO     `json:"image,omitempty"`
	Copyright   string       `json:"copyRight,omitempty"`
	Articles    []ArticleDTO `json:"articles,omitempty"`
	FeedType    string       `json:"feedType,omitempty"`
	Slug        string       `json:"slug,omitempty"`
}

func (n NewsfeedDTO) ToDomain() entity.Newsfeed {
	articles := make([]entity.Article, 0)
	for _, a := range n.Articles {
		articles = append(articles, a.ToDomain())
	}
	newsfeed := entity.Newsfeed{
		ID:          n.ID,
		Title:       n.Title,
		Description: n.Description,
		FeedURL:     n.FeedURL,
		SiteURL:     n.SiteURL,
		Language:    n.Language,
		Image:       n.Image.ToDomain(),
		Copyright:   n.Copyright,
		Articles:    articles,
		FeedType:    n.FeedType,
		Slug:        n.Slug,
	}
	return newsfeed
}
