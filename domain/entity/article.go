package entity

import (
	"encoding/json"

	"newser.app/domain/value"
)

type Article struct {
	Item
	Read  bool `json:"read"`
	Saved bool `json:"saved"`
}

type Item struct {
	ID          ID         `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Content     string     `json:"content"`
	Link        value.Link `json:"link"`
	Updated     string     `json:"updated"`
	Published   string     `json:"published"`
	Author      *Person    `json:"author"`
	GUID        string     `json:"guid"`
	Image       *Image     `json:"image"`
	Categories  []string   `json:"categories"`
}

func (a Article) JSON() []byte {
	j, _ := json.MarshalIndent(a, "", "  ")
	return j
}

func (a Article) String() string {
	return string(a.JSON())
}

func (a *Article) SetRead() {
	a.Read = true
}

func (a *Article) SetSaved() {
	a.Saved = true
}

func (a *Article) SetUnread() {
	a.Read = false
}
