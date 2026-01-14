package xurl

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractWords(t *testing.T) {
	rawURLs := []string{
		"https://example.com/api/core/?query=param&query=param2&query2=param3#fragment",
		"https://another.com/path/to/resource?foo=bar&baz=qux",
		"https://example.com/api/core/?query=param",
	}

	words, errs := ExtractWords(rawURLs)

	for _, e := range errs {
		require.NoError(t, e)
	}

	require.NotEmpty(t, words, "expected some words got none")

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
