package main

import (
	"testing"
)

func TestGetMD5Hash(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "Normal URL",
			text: "www.example.com",
			want: "7c1767b30512b6003fd3c2e618a86522",
		},
		{
			name: "Empty URL",
			text: "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashText := GetMD5Hash(tt.text)
			if tt.want != hashText {
				t.Errorf("want:%s got: %s", tt.want, hashText)
			}
		})
	}
}
