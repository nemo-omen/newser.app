package entity

import "time"

type Highlight struct {
	ID           ID        `json:"id"`
	StartOffset  int       `json:"startOffset"`
	EndOffset    int       `json:"endOffset"`
	UserId       ID        `json:"userId"`
	AnnotationID ID        `json:"annotationId"`
	ArticleID    ID        `json:"articleId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
