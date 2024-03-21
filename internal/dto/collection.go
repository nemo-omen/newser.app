package dto

type CollectionDTO struct {
	ID     string `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`
	Title  string `json:"title" db:"title"`
	Slug   string `json:"slug" db:"slug"`
}
