package repository

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strings"

	"github.com/weppos/publicsuffix-go/publicsuffix"
	"github.com/zomasec/tld"
	"newser.app/pkg/constant"
)

var (
	ErrInvalidUrl = errors.New("not a valid url")
)

func GetValidUrl(u string) (string, error) {
	if IsFeedLink(u) {
		return u, nil
	}

	if !HasValidTLD(u) {
		return "", ErrInvalidUrl
	}

	parsed, err := ParseURL(u)
	if err != nil {
		return "", fmt.Errorf("error parsing as url: %w", err)
	}

	if !IsSite(parsed.String()) {
		return "", ErrInvalidUrl
	}
	return parsed.String(), nil
}

func HasValidTLD(u string) bool {
	_, err := publicsuffix.Parse(u)
	return err == nil
}

func ParseURL(u string) (*url.URL, error) {
	parsed, err := tld.Parse(u)
	if err != nil {
		return nil, fmt.Errorf("error parsing URL: %w", err)
	}
	if parsed.Scheme == "" {
		parsed.Scheme = "https"
	}
	return parsed.URL, nil
}

func IsSite(url string) bool {
	client := http.Client{}
	res, err := client.Head(url)
	if err != nil {
		return false
	}
	if res.StatusCode == 200 {
		return true
	}
	return false
}

func HasCommonFeedPath(u string) bool {
	for _, path := range constant.CommonFeedPaths {
		if strings.HasSuffix(u, path) {
			return true
		}
	}
	return false
}

func HasCommonFeedExtension(u string) bool {
	for _, ext := range constant.CommonFeedExtensions {
		if strings.HasSuffix(u, ext) {
			return true
		}
	}
	return false
}

func IsFeedLink(u string) bool {
	res, err := http.Get(u)
	if err != nil {
		return false
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type")
	return slices.Contains(constant.ValidDocContentTypes, constant.DocContentType(contentType))
}
