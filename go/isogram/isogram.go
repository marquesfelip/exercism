package isogram

import (
	"sort"
	"strings"
)

func IsIsogram(word string) bool {
	runes := []rune(strings.ToLower(word))
	var charCheck rune

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	for _, char := range runes {
		if string(char) == " " || string(char) == "-" {
			continue
		}

		if charCheck == char {
			return false
		}
		charCheck = char
	}

	return true
}
