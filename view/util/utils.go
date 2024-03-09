package util

import (
	"context"
	"strconv"
)

func GetImgAlt(imgTitle, feedTitle string) string {
	alt := ""
	if imgTitle != "" {
		alt = imgTitle
	} else if feedTitle != "" {
		alt = feedTitle
	}
	return alt
}

func IdToString(id int64) string {
	return strconv.FormatInt(id, 10)
}

func GetPageTitle(ctx context.Context) string {
	title := ctx.Value("title")
	if title != nil {
		return title.(string)
	}
	return ""
}

func GetUserViewPreference(ctx context.Context) string {
	viewPref := ctx.Value("viewPreference")
	if viewPref != nil {
		return viewPref.(string)
	}
	return ""
}
