package entity

type Newsfeed struct {
	ID          int64
	Title       string
	Description string
	FeedURL     string
	SiteURL     string
	Language    string
	Image       Image
	Copyright   string
	Articles    []Article
	FeedType    string
	Slug        string
}
