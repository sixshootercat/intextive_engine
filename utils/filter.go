package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball"
)

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

// stopwordFilter removes stopwords from the given list of tokens.
// Stopwords are common words that are often considered insignificant
// and are typically filtered out in text processing tasks.
// This function takes a list of tokens and returns a new list with
// the stopwords removed.
func stopwordFilter(tokens []string) []string {
	stopword := map[string]struct{}{
		"a":    {},
		"and":  {},
		"the":  {},
		"in":   {},
		"of":   {},
		"to":   {},
		"that": {},
		"i":    {},
	}

	r := make([]string, 0, len(tokens))

	for _, token := range tokens {
		if _, ok := stopword[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		stemmedToken, _ := snowballeng.Stem(token, "english", false)
		r[i] = stemmedToken
	}
	return r
}
