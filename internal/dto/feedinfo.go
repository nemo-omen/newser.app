package dto

type FeedInfoDTO struct {
	FeedId         string `json:"feedId"`
	FeedTitle      string `json:"feedTitle"`
	SubscriptionId string `json:"subscriptionId"`
	ImageUrl       string `json:"imageUrl"`
	ImageTitle     string `json:"imageTitle"`
	UnreadCount    int    `json:"unreadCount"`
}

type FeedInfoDAO struct {
	FeedId         string
	FeedTitle      string
	SubscriptionId string
	ImageUrl       string
	ImageTitle     string
	UnreadCount    int
}

func NewFeedInfoDTO(feedId, feedTitle, subscriptionId, imageUrl, imageTitle string, unreadCount int) *FeedInfoDTO {
	return &FeedInfoDTO{
		FeedId:         feedId,
		FeedTitle:      feedTitle,
		SubscriptionId: subscriptionId,
		ImageUrl:       imageUrl,
		ImageTitle:     imageTitle,
		UnreadCount:    unreadCount,
	}
}

func NewFeedInfoDAO(feedId, feedTitle, subscriptionId, imageUrl, imageTitle string, unreadCount int) *FeedInfoDAO {
	return &FeedInfoDAO{
		FeedId:         feedId,
		FeedTitle:      feedTitle,
		SubscriptionId: subscriptionId,
		ImageUrl:       imageUrl,
		ImageTitle:     imageTitle,
		UnreadCount:    unreadCount,
	}
}
