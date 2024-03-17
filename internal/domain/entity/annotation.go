package entity

import (
	"time"
)

type Annotation struct {
	ID        ID        `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserId    ID        `json:"userId"`
	ArticleId ID        `json:"articleId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
