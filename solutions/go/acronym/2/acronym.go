package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	initialWord := true
	acronym := make([]rune, 0)

	for _, v := range s {
		if v == ' ' || v == '-' {
			initialWord = true
		}

		if initialWord {
			if unicode.IsLetter(v) {
				acronym = append(acronym, v)
				initialWord = false
			}
		}
	}

	return strings.ToUpper(string(acronym))
}
