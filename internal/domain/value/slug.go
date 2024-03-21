package value

import (
	"newser.app/shared/util"
)

type Slug string

func NewSlug(s string) (Slug, error) {
	slug := util.Slugify(s)
	return Slug(slug), nil
}

func (s Slug) String() string {
	return string(s)
}
