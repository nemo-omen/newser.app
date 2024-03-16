package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonEntity(t *testing.T) {
	testCases := []struct {
		name  string
		email string
		err   error
	}{
		{
			name:  "John Doe",
			email: "john@doe.whatever",
			err:   nil,
		},
		{
			name:  "",
			email: "john@doe.whatever",
			err:   ErrInvalidInput,
		},
		{
			name:  "John Doe",
			email: "",
			err:   ErrInvalidInput,
		},
		{
			name:  "John Doe",
			email: "@doe.whatever",
			err:   ErrInvalidInput,
		},
		{
			name:  "John Doe",
			email: "whatever",
			err:   ErrInvalidInput,
		},
	}

	for _, tc := range testCases {
		_, err := NewPerson(tc.name, tc.email)
		assert.Equal(t, tc.err, err)
	}
}
