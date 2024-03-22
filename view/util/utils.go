package util

import (
	"context"
	"fmt"
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

// GetPageTitle returns the title of the page
// from the context. This works on page load.
// HTMX is used on subsequent navigation
// changes to replace the title on the page
func GetPageTitle(ctx context.Context) string {
	title, ok := ctx.Value("title").(string)
	if !ok {
		return ""
	}
	fmt.Println("title: ", title)
	return title
}

// "expanded" or "collapsed"
func GetLayoutPreference(ctx context.Context) string {
	layoutPref := ctx.Value("layout")
	if layoutPref != nil {
		return layoutPref.(string)
	}
	return "expanded"
}

// "read" or "unread
func GetViewPreference(ctx context.Context) string {
	viewPref, ok := ctx.Value("view").(string)
	if ok {
		return viewPref
	}
	return "unread"
}
