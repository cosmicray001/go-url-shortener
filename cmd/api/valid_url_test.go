package main

import (
	"testing"
)

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		inputUrl string
		want     bool
	}{
		// Valid URLs with 2xx status codes
		{"Valid URL", "https://www.google.com", true},
		{"Valid URL", "https://github.com/cosmicray001", true},
		{"Valid URL", "https://www.example.com", true},

		// Invalid URLs with non-2xx status codes
		{"Invalid URL", "https://www.nonexistent-link.com", false},
		{"Invalid URL", "https://www.invalid-link.com", false},

		// Empty inputURL should return false
		{"empty URL", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidURL(tt.inputUrl)
			if tt.want != result {
				t.Errorf("want:%v got: %v", tt.want, result)
			}
		})
	}
}
