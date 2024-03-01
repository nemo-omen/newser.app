package util

import "github.com/gosimple/slug"

func Slugify(str string) string {
	return slug.Make(str)
}
