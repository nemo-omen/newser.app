package entity

import "encoding/json"

type SubscriptionFeedCreate struct {
	Feed   Newsfeed
	UserID int64
}

func (s SubscriptionFeedCreate) JSON() []byte {
	j, _ := json.MarshalIndent(s, "", "  ")
	return j
}

func (s SubscriptionFeedCreate) String() string {
	return string(s.JSON())
}
