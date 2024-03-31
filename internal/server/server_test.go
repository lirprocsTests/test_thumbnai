package server

import (
	"testing"
)

func TestGetVideoID(t *testing.T) {
	url := "https://www.youtube.com/watch?v=abc123"
	expected := "abc123"
	result := getVideoID(url)
	if result != expected {
		t.Errorf("getVideoID(%s) = %s; want %s", url, result, expected)
	}
}
