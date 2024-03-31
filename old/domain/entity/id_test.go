package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewID(t *testing.T) {
	id := NewID()
	assert.NotEqual(t, ID{}, id)
}

func TestNewIDFromString(t *testing.T) {
	testCases := []struct {
		input     string
		expectErr bool
	}{
		{
			input:     "c4a760a8-dbcf-5254-a0d9-6a4474bd1b62",
			expectErr: false,
		},
		{
			input:     "invalid-uuid",
			expectErr: true,
		},
	}

	for _, tc := range testCases {
		id, err := NewIDFromString(tc.input)
		if tc.expectErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.NotEqual(t, ID{}, id)
		}
	}
}
