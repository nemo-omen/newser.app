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

// GetValidUrl attempts to transform and validate
// a given string into a valid URL.
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

// HasValidTLD checks if a given string has a valid TLD.
func HasValidTLD(u string) bool {
	_, err := publicsuffix.Parse(u)
	return err == nil
}

// ParseURL parses a given string as a URL.
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

// IsSite checks if a given URL is a valid site
// by sending a HEAD request and checking if the
// response StatusCode == 200.
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

// HasCommonFeedPath checks if a given URL has a common
// feed path suffix.
func HasCommonFeedPath(u string) bool {
	for _, path := range constant.CommonFeedPaths {
		if strings.HasSuffix(u, path) {
			return true
		}
	}
	return false
}

// HasCommonFeedExtension checks if a given URL has a common
// feed extension suffix.
func HasCommonFeedExtension(u string) bool {
	for _, ext := range constant.CommonFeedExtensions {
		if strings.HasSuffix(u, ext) {
			return true
		}
	}
	return false
}

// IsFeedLink checks if a given URL is a valid feed link
// by sending a GET request and checking if the response
// Content-Type is a valid feed type.
func IsFeedLink(u string) bool {
	res, err := http.Get(u)
	if err != nil {
		return false
	}
	defer res.Body.Close()

	contentType := res.Header.Get("Content-Type")
	return slices.Contains(constant.ValidDocContentTypes, constant.DocContentType(contentType))
}
