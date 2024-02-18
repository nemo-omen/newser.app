package util

import (
	"log"
	"net/url"
	"strings"
)

func IsUrl(str string) bool {
	u, err := url.ParseRequestURI(str)
	if err != nil {
		log.Printf("err=%+v, url=%+v\n", err, u)
		return false
	}

	if !strings.Contains(u.Host, ".") {
		return false
	}

	return true
}
