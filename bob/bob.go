// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "regexp"

// isQuestion returns true if the string ends with a question mark and spaces
func isQuestion(s string) bool {
	reg := regexp.MustCompile(`\?[[:space:]]*$`)
	return reg.MatchString(s)
}

// isAllSpaces returns true if the string is only made of spaces
func isAllSpaces(s string) bool {
	reg := regexp.MustCompile(`^[[:space:]]*$`)
	return reg.MatchString(s)
}

// isAllCapitals returns true if the string is only made of capital letters (and spaces)
func isAllCapitals(s string) bool {
	regCaps := regexp.MustCompile(`[a-z+]`)
	regLow := regexp.MustCompile(`[A-Z+]`)
	return !regCaps.MatchString(s) && regLow.MatchString(s)
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if isAllSpaces(remark) {
		return "Fine. Be that way!"
	} else if isAllCapitals(remark) {
		if isQuestion(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} else if isQuestion(remark) {
		return "Sure."
	}
	return "Whatever."
}
