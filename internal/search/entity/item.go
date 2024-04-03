package entity

import (
	"encoding/json"
	"time"
)

type Item struct {
	Title           string
	Description     string
	Content         string
	Link            string
	Links           []string
	Updated         string
	UpdatedParsed   *time.Time
	Published       string
	PublishedParsed *time.Time
	Authors         []*Person
	GUID            string
	Image           *Image
	Categories      []string
}

func (i *Item) JSON() []byte {
	j, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return []byte{}
	}
	return j
}

func (i *Item) String() string {
	return string(i.JSON())
}
