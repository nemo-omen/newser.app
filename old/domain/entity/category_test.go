package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"newser.app/internal/domain/value"
)

func TestNewCategory(t *testing.T) {
	testCases := []struct {
		term       string
		expectNil  bool
		expectTerm value.Term
	}{
		{
			term:       "example",
			expectNil:  false,
			expectTerm: value.Term("example"),
		},
		{
			term:       "",
			expectNil:  true,
			expectTerm: "",
		},
	}

	for _, tc := range testCases {
		category := NewCategory(tc.term)
		if tc.expectNil {
			assert.Nil(t, category)
		} else {
			assert.NotNil(t, category)
			assert.Equal(t, tc.expectTerm, category.Term)
		}
	}
}

func TestCategory_JSON(t *testing.T) {
	// Test data
	category := Category{
		ID:   NewID(),
		Term: value.Term("example"),
	}

	// Call the JSON method
	jsonData := category.JSON()

	// Unmarshal the JSON data
	var decodedCategory Category
	err := json.Unmarshal(jsonData, &decodedCategory)
	assert.NoError(t, err)

	// Assert the values
	assert.Equal(t, category.ID, decodedCategory.ID)
	assert.Equal(t, category.Term, decodedCategory.Term)
}
