package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTerm(t *testing.T) {
	// Test valid term
	term, err := NewTerm("example")
	assert.NoError(t, err)
	assert.Equal(t, Term("example"), term)

	// Test empty term
	term, err = NewTerm("")
	assert.Error(t, err)
	assert.Equal(t, Term(""), term)
}

func TestTermString(t *testing.T) {
	term := Term("example")
	assert.Equal(t, "example", term.String())
}
