package dto

type FeedInfoDTO struct {
	FeedId         string `json:"feedId" db:"feedId"`
	FeedTitle      string `json:"feedTitle" db:"feedTitle"`
	SubscriptionId string `json:"subscriptionId" db:"subscriptionId"`
	ImageUrl       string `json:"imageUrl" db:"imageUrl"`
	ImageTitle     string `json:"imageTitle" db:"imageTitle"`
	UnreadCount    int    `json:"unreadCount" db:"unreadCount"`
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
