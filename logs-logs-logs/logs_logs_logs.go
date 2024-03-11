package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		switch char {
		case 0x2757:
			return "recommendation"
		case 0x1F50d:
			return "search"
		case 0x2600:
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	s := strings.Builder{}
	s.Grow(len(log))
	for _, char := range log {
		if char == oldRune {
			s.WriteRune(newRune)
		} else {
			s.WriteRune(char)
		}
	}
	return s.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
