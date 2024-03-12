package reverse

import (
	"strings"
)

func Reverse(input string) string {
	if input == "" {
		return input
	}

	runes := make([]rune, 0, len(input))
	for _, r := range input {
		runes = append(runes, r)
	}

	sb := strings.Builder{}
	sb.Grow(len(runes))
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}
	return sb.String()
}
