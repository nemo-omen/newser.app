package entity

import (
	"encoding/json"
	"time"
)

type Feed struct {
	Title           string
	Description     string
	Link            string
	FeedLink        string
	Links           []string
	Updated         string
	UpdatedParsed   *time.Time
	Published       string
	PublishedParsed *time.Time
	Authors         []*Person
	Language        string
	Image           *Image
	Copyright       string
	Categories      []string
	Items           []*Item
	FeedType        string
}

func (f *Feed) JSON() []byte {
	j, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return []byte{}
	}
	return j
}

func (f *Feed) String() string {
	return string(f.JSON())
}
