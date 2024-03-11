package annotation

import "newser.app/internal/domain/article"

type Annotation struct {
	note *Note
	item *article.Item
}
