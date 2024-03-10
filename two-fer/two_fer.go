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

// ShareWith2 returns a "One for X, one for me" string.
// X is either the provided name or the default string "you".
// This version does not use Sprintf, in order to benchmark.
// This version is much faster.
func ShareWith2(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
