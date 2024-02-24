package util

import (
	"testing"
)

func TestIsUrl(t *testing.T) {
	tests := []struct {
		Desc     string
		Input    string
		Expected bool
	}{
		{"Test empty string", "", false},
		{"Test url with no scheme & no domain", "jeffcaldwell", false},
		{"Test url with scheme and no domain", "https://jeffcaldwell", false},
		{"Test url with domain and no scheme", "jeffcaldwell.is", false},
		{"Test full url", "https://jeffcaldwell.is", true},
	}

	for _, tt := range tests {
		actual := IsUrl(tt.Input)
		expected := tt.Expected
		if actual != expected {
			t.Errorf("expected: %v, actual: %v", expected, actual)
		}
	}
}
