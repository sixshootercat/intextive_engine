package utils

import (
	"strings"
	"unicode"
)

// tokenize splits a string into tokens
func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

// analyze takes a string of text and performs tokenization, lowercase filtering,
// stopword filtering, and stemming on it. It returns a slice of strings
// representing the analyzed tokens.
func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopwordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
