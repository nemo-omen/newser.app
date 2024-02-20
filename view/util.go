package view

import (
	"context"
	"strconv"
)

func GetLen[T any](s []T) string {
	return strconv.Itoa(len(s))
}

func GetIsHx(ctx context.Context) bool {
	isHx := ctx.Value("isHx")
	if isHx != nil {
		return isHx.(bool)
	}
	return false
}
