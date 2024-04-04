package dto

import (
	"encoding/json"
)

type FeedInfoDTO struct {
	FeedId      string `json:"feedId"`
	FeedTitle   string `json:"feedTitle"`
	ImageUrl    string `json:"imageUrl"`
	ImageTitle  string `json:"imageTitle"`
	UnreadCount int    `json:"unreadCount"`
}

func (f FeedInfoDTO) JSON() []byte {
	j, _ := json.MarshalIndent(f, "", "  ")
	return j
}

func (f FeedInfoDTO) String() string {
	return string(f.JSON())
}
