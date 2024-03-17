package value

import (
	"net/url"

	"newser.app/shared"
)

// Link represents a URL.
type Link string

// NewLink creates a new Link.
func NewLink(link string) (Link, error) {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		appErr := shared.NewAppError(
			ErrInvalidInput,
			"Not a valid URL",
			"NewLink",
			"value.Link",
		)
		return "", appErr
	}
	return Link(link), nil
}

// String returns the string representation of a Link.
func (l Link) String() string {
	return string(l)
}
