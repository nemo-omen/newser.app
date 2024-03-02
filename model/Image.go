package model

type Image struct {
	ID    int64  `json:"id,omitempty" db:"id"`
	URL   string `json:"url,omitempty" db:"url"`
	Title string `json:"title,omitempty" db:"title"`
}
