package atbash

import (
	"bytes"
)

func Atbash(s string) string {
	output := make([]byte, 0, len(s)+len(s)/5)
	for _, chr := range []byte(s) {
		if '0' <= chr && chr <= '9' {
			output = append(output, chr)
		} else if 'a' <= chr && chr <= 'z' {
			output = append(output, 'z'-chr+'a')
		} else if 'A' <= chr && chr <= 'Z' {
			output = append(output, 'Z'-chr+'a')
		} else {
			continue
		}
		// Add a space every 5 characters, use a hack to account for the spaces we add.
		if (len(output)+1)%6 == 0 {
			output = append(output, ' ')
		}
	}
	return string(bytes.TrimSuffix(output, []byte{' '}))
}
