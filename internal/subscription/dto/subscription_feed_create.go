package dto

import (
	"encoding/json"
)

type SubscriptionFeedCreateDTO struct {
	Feed   NewsfeedDTO `json:"feed"`
	UserID int64       `json:"user_id"`
}

func (s SubscriptionFeedCreateDTO) JSON() []byte {
	j, _ := json.MarshalIndent(s, "", "  ")
	return j
}

func (s SubscriptionFeedCreateDTO) String() string {
	return string(s.JSON())
}
