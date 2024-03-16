package entity

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"newser.app/domain/value"
)

func TestNewImage(t *testing.T) {
	testCases := []struct {
		url   string
		title string
	}{
		{
			url:   "https://example.com/image.jpg",
			title: "Example Corp.",
		},
		{
			url:   "",
			title: "Example Corp.",
		},
	}

	for _, tc := range testCases {
		image := NewImage(tc.url, tc.title)
		if tc.url == "" {
			assert.Nil(t, image)
		} else {
			assert.NotNil(t, image)
			assert.Equal(t, value.Link(tc.url), image.URL)
			assert.Equal(t, tc.title, image.Title)
		}
	}
}

func TestImage_JSON(t *testing.T) {
	// Test data
	image := Image{
		URL:   "https://example.com/image.jpg",
		Title: "Example Corp.",
	}

	// Call the JSON method
	jsonData := image.JSON()

	// Unmarshal the JSON data
	var decodedImage Image
	err := json.Unmarshal(jsonData, &decodedImage)
	assert.NoError(t, err)

	// Assert the values
	assert.Equal(t, image.URL, decodedImage.URL)
	assert.Equal(t, image.Title, decodedImage.Title)
}
