package xurl

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

func ExtractWords(rawURLs []string) ([]string, []error) {
	seen := make(map[string]struct{})
	var wordlist []string
	var errs []error

	for _, rawURL := range rawURLs {
		words, err := ExtractURL(rawURL)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		for _, word := range words {
			if _, exists := seen[word]; !exists {
				seen[word] = struct{}{}
				wordlist = append(wordlist, word)
			}
		}
	}

	return wordlist, errs
}

func ExtractURL(rawURL string) ([]string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to Parse URL %s: %v", rawURL, err)
	}

	var words []string

	// append url paths to words
	paths := strings.SplitSeq(parsedURL.Path, "/")
	for p := range paths {
		p = strings.TrimSpace(p)
		if !slices.Contains(words, p) {
			words = append(words, p)
		}
	}

	// append url query params to words
	queryParams := parsedURL.Query()
	for query, params := range queryParams {
		query = strings.TrimSpace(query)
		if !slices.Contains(words, query) {
			words = append(words, query)
		}
		for _, param := range params {
			param = strings.TrimSpace(param)
			if !slices.Contains(words, param) {
				words = append(words, param)
			}
		}
	}

	// append url fragment
	fragment := strings.TrimSpace(parsedURL.Fragment)
	if !slices.Contains(words, fragment) {
		words = append(words, fragment)
	}

	return words, nil
}
