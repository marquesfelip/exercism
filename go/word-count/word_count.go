package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

var rePattern = regexp.MustCompile(`[a-zA-Z0-9]+('[a-zA-Z0-9]+)?`)

func WordCount(phrase string) Frequency {
	wordCount := make(Frequency)
	words := rePattern.FindAllString(phrase, -1)

	for _, v := range words {
		v := strings.ToLower(v)
		wordCount[v]++
	}

	return wordCount
}
