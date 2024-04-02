package dto

type FeedInfoDTO struct {
	FeedId      string `json:"feedId" db:"feed_id"`
	FeedTitle   string `json:"feedTitle" db:"feed_title"`
	ImageUrl    string `json:"imageUrl" db:"image_url"`
	ImageTitle  string `json:"imageTitle" db:"image_title"`
	UnreadCount int    `json:"unreadCount" db:"unread_count"`
}
