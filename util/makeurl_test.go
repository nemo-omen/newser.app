package util

import (
	"net/url"
	"testing"
)

func TestMakeUrl(t *testing.T) {
	validURL, _ := url.ParseRequestURI("https://jeffcaldwell.is")
	tests := []struct {
		desc     string
		input    string
		expected *url.URL
	}{
		{"Empty string", "", &url.URL{}},
		{"No scheme or host", "jeffcaldwell", &url.URL{}},
		{"No host", "https://", &url.URL{}},
		{"No scheme", "jeffcaldwell.is", validURL},
		{"Valid URL string", "https://jeffcaldwell.is", validURL},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			expected := tt.expected
			actual, _ := MakeUrl(tt.input)

			if expected.Host != actual.Host || expected.Scheme != actual.Scheme {
				t.Errorf("input %+v, expected %+v, actual %+v", tt.input, expected, actual)
			}
		})
	}
}
