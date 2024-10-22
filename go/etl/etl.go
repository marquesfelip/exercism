package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	letterMapping := map[string]int{}

	for value, letters := range in {
		for _, letter := range letters {
			letterMapping[strings.ToLower(letter)] = value
		}
	}

	return letterMapping
}
