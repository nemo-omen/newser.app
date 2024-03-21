package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLink_ValidLink(t *testing.T) {
	link := "https://example.com"
	l, err := NewLink(link)
	assert.NoError(t, err)
	assert.Equal(t, Link(link), l)
}

func TestNewLink_InvalidLink(t *testing.T) {
	link := "invalid-url"
	_, err := NewLink(link)
	assert.Error(t, err)
	assert.EqualError(t, err, ErrInvalidInput.Error())
}

func TestLink_String(t *testing.T) {
	link := "https://example.com"
	l := Link(link)
	assert.Equal(t, link, l.String())
}
