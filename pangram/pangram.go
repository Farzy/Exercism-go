package pangram

import (
	"unicode"
)

func IsPangram(input string) bool {
	var bitFlags uint32
	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		bit := int(unicode.ToLower(r) - 'a')
		bitFlags |= 1 << bit
	}
	return bitFlags == 1<<26-1
}
