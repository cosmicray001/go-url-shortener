package main

import (
	"net/http"
)

func isValidURL(inputURL string) bool {
	if inputURL == "" {
		return false
	}

	resp, err := http.Head(inputURL)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true
	}

	return false
}
