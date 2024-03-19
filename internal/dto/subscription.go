package dto

type SubscriptionDTO struct {
	UserID string `json:"user_id" db:"user_id"`
	FeedID string `json:"feed_id" db:"feed_id"`
}
