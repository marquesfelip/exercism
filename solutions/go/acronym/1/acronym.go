package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})

	var acronym strings.Builder

	for _, word := range words {
		word = removeNonAlhpa(word)
		acronym.WriteByte(word[0])
	}

	return strings.ToUpper(acronym.String())
}

func removeNonAlhpa(s string) string {
	var builder strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) {
			builder.WriteRune(r)
		}
	}

	return builder.String()
}
