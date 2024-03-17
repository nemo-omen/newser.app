package value

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail(t *testing.T) {
	testCases := []struct {
		input     string
		expectErr bool
	}{
		{
			input:     "test@example.com",
			expectErr: false,
		},
		{
			input:     "invalid-email",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		email, err := NewEmail(tc.input)
		if tc.expectErr {
			assert.EqualError(t, err, ErrInvalidInput.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tc.input, email.String())
		}
	}
}

func TestEmailString(t *testing.T) {
	email := Email("test@example.com")
	assert.Equal(t, "test@example.com", email.String())
}
