// Package twofer implements "two for one"
package twofer

import "fmt"

// ShareWith returns a "One for X, one for me" string.
// X is either the provided name or the default string "you".
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
