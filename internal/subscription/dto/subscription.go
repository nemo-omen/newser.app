package dto

import (
	"encoding/json"

	"newser.app/internal/subscription/entity"
)

type SubscriptionDTO struct {
	UserID string `json:"user_id"`
	FeedID string `json:"feed_id"`
}

type CreateSubscriptionRequestDTO struct {
	FeedID int64 `json:"feedId"`
	UserID int64 `json:"userId"`
}

func (s SubscriptionDTO) JSON() []byte {
	j, _ := json.MarshalIndent(s, "", "  ")
	return j
}

func (s SubscriptionDTO) String() string {
	return string(s.JSON())
}

func (s SubscriptionDTO) ToDomain() entity.Subscription {
	return entity.Subscription{
		UserID: s.UserID,
		FeedID: s.FeedID,
	}
}

func (r CreateSubscriptionRequestDTO) JSON() []byte {
	j, _ := json.MarshalIndent(r, "", "  ")
	return j
}

func (r CreateSubscriptionRequestDTO) String() string {
	return string(r.JSON())
}

func (r CreateSubscriptionRequestDTO) ToDomain() entity.CreateSubscriptionRequest {
	return entity.CreateSubscriptionRequest{
		FeedID: r.FeedID,
		UserID: r.UserID,
	}
}
