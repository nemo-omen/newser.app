package value

import (
	"net/url"
)

// Link represents a URL.
type Link string

// NewLink creates a new Link.
func NewLink(link string) (Link, error) {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return "", ErrInvalidInput
	}
	return Link(link), nil
}

// String returns the string representation of a Link.
func (l Link) String() string {
	return string(l)
}
