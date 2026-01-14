package xurl

import (
	"testing"
)

func TestExtractWords(t *testing.T) {
	rawURLs := []string{
		"https://example.com/api/core/?query=param&query=param2&query2=param3#fragment",
		"https://another.com/path/to/resource?foo=bar&baz=qux",
		"https://example.com/api/core/?query=param",
	}

	words, errs := ExtractWords(rawURLs)

	if len(errs) > 0 {
		for _, err := range errs {
			t.Errorf("error extracting words: %v", err)
		}
	}

	if len(words) == 0 {
		t.Errorf("expected some words, got none")
	}

	for _, word := range words {
		t.Log(word)
	}
}

func TestExtractURL(t *testing.T) {
	rawURL := "https://example.com/api/core/?query=param&query=param2&query2=param3#fragment"
	words, err := ExtractURL(rawURL)
	if err != nil {
		t.Errorf("error extracting words from url %s: %v", rawURL, err)
	}
	for _, word := range words {
		t.Log(word)
	}
}
