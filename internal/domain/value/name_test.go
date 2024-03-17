package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewName(t *testing.T) {
	testCases := []struct {
		input     string
		expectErr bool
	}{
		{
			input:     "John",
			expectErr: false,
		},
		{
			input:     "",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		name, err := NewName(tc.input)
		if tc.expectErr {
			assert.EqualError(t, err, ErrInvalidInput.Error())
			assert.Equal(t, Name(""), name)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, Name(tc.input), name)
		}
	}
}

func TestNameString(t *testing.T) {
	name := Name("John")
	assert.Equal(t, "John", name.String())
}
