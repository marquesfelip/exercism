// Package twofer provides a function to generate a string in the format
// "One for X, one for me." where X is a given name. If no name is provided,
// it defaults to "you".
package twofer

import "fmt"

// ShareWith returns a string in the format "One for X, one for me."
// where X is the provided name. If no name is provided, it defaults to "you".
func ShareWith(name string) string {
	if name != "" {
		return fmt.Sprintf("One for %s, one for me.", name)
	}
	return "One for you, one for me."
}
