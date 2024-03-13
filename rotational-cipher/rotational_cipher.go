package rotationalcipher

import (
	"strings"
)

func RotationalCipher(plain string, shiftKey int) string {
	var out strings.Builder
	out.Grow(len(plain))

	for _, chr := range plain {
		if 'a' <= chr && chr <= 'z' {
			chr = 'a' + (chr-'a'+rune(shiftKey))%26
		}
		if 'A' <= chr && chr <= 'Z' {
			chr = 'A' + (chr-'A'+rune(shiftKey))%26
		}
		out.WriteRune(chr)
	}
	return out.String()
}
