package view

import "strconv"

func GetLen[T any](s []T) string {
	return strconv.Itoa(len(s))
}
