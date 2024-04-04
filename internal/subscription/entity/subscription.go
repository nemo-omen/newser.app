package entity

import (
	"encoding/json"
)

type Subscription struct {
	ID     int64
	UserID string
	FeedID string
}

type CreateSubscriptionRequest struct {
	FeedID int64
	UserID int64
}

func (s Subscription) JSON() []byte {
	j, _ := json.MarshalIndent(s, "", "  ")
	return j
}

func (s Subscription) String() string {
	return string(s.JSON())
}

func (r CreateSubscriptionRequest) JSON() []byte {
	json, _ := json.MarshalIndent(r, "", "  ")
	return json
}

func (r CreateSubscriptionRequest) String() string {
	return string(r.JSON())
}
