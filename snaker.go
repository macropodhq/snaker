// Package snaker provides methods to convert CamelCase names to snake_case and back.
// It considers the list of allowed initialsms used by github.com/golang/lint/golint (e.g. ID or HTTP)
package snaker

import (
	"strings"
	"unicode"
)

func Split(s string) []string {
	if strings.Contains(s, "_") {
		return SplitSnake(s)
	} else {
		return SplitCamel(s)
	}
}

func SplitCamel(s string) []string {
	var words []string
	var lastPos int
	rs := []rune(s)

	for i := 0; i < len(rs); i++ {
		if i > 0 && unicode.IsUpper(rs[i]) {
			if initialism := startsWithInitialism(s[lastPos:]); initialism != "" {
				words = append(words, initialism)

				i += len(initialism) - 1
				lastPos = i
				continue
			}

			words = append(words, s[lastPos:i])
			lastPos = i
		}
	}

	// append the last word
	if s[lastPos:] != "" {
		words = append(words, s[lastPos:])
	}

	return words
}

func SplitSnake(s string) []string {
	return strings.Split(s, "_")
}

func ToCamel(s string, lower bool) string {
	var result string

	words := Split(s)

	for i, word := range words {
		if upper := strings.ToUpper(word); commonInitialisms[upper] {
			if lower && i == 0 {
				upper = strings.ToLower(upper)
			}

			result += upper

			continue
		}

		if lower && i == 0 {
			word = strings.ToLower(word)
		} else {
			w := []rune(word)
			w[0] = unicode.ToUpper(w[0])
			word = string(w)
		}

		result += word
	}

	return result
}

func ToLowerCamel(s string) string {
	return ToCamel(s, true)
}

func ToUpperCamel(s string) string {
	return ToCamel(s, false)
}

func ToSnake(s string) string {
	var result string

	words := Split(s)

	for i, word := range words {
		if i > 0 {
			result += "_"
		}

		result += strings.ToLower(word)
	}

	return result
}

// startsWithInitialism returns the initialism if the given string begins with it
func startsWithInitialism(s string) string {
	var initialism string
	// the longest initialism is 5 char, the shortest 2
	for i := 1; i <= 5; i++ {
		if len(s) > i-1 && commonInitialisms[s[:i]] {
			initialism = s[:i]
		}
	}
	return initialism
}

// commonInitialisms, taken from
// https://github.com/golang/lint/blob/3d26dc39376c307203d3a221bada26816b3073cf/lint.go#L482
var commonInitialisms = map[string]bool{
	"API":   true,
	"ASCII": true,
	"CPU":   true,
	"CSS":   true,
	"DNS":   true,
	"EOF":   true,
	"GUID":  true,
	"HTML":  true,
	"HTTP":  true,
	"HTTPS": true,
	"ID":    true,
	"IP":    true,
	"JSON":  true,
	"LHS":   true,
	"QPS":   true,
	"RAM":   true,
	"RHS":   true,
	"RPC":   true,
	"SLA":   true,
	"SMTP":  true,
	"SSH":   true,
	"TLS":   true,
	"TTL":   true,
	"UI":    true,
	"UID":   true,
	"UUID":  true,
	"URI":   true,
	"URL":   true,
	"UTF8":  true,
	"VM":    true,
	"XML":   true,
}
