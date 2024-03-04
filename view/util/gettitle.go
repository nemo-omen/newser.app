package util

import "context"

func GetPageTitle(ctx context.Context) string {
	title := ctx.Value("title")
	if title != nil {
		return title.(string)
	}
	return ""
}
