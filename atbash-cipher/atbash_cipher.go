package atbash

import (
	"strings"
	"unicode"
)

const atbashAlphabet = "zyxwvutsrqponmlkjihgfedcba"

func Atbash(s string) string {
	output := strings.Builder{}
	for _, chr := range s {
		if unicode.IsNumber(chr) {
			output.WriteRune(chr)
		} else if !unicode.IsLetter(chr) {
			continue
		} else {
			output.WriteByte(atbashAlphabet[unicode.ToLower(chr)-'a'])
		}
		// Add a space every 5 characters, use a hack to account for the spaces we add.
		if (output.Len()+1)%6 == 0 {
			output.WriteByte(' ')
		}
	}
	return strings.TrimSpace(output.String())
}
