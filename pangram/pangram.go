package pangram

/*
goos: darwin
goarch: amd64
pkg: pangram
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkPangram
BenchmarkPangram-12           	 3092238	       361.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkPangramUnicode
BenchmarkPangramUnicode-12    	  945254	      1225 ns/op	       0 B/op	       0 allocs/op
*/

import (
	"unicode"
)

func IsPangram(input string) bool {
	var bitFlags uint32
	for _, r := range input {
		if 'a' <= r && r <= 'z' {
			bitFlags |= 1 << int(r-'a')
		} else if 'A' <= r && r <= 'Z' {
			bitFlags |= 1 << int(r-'A')
		}
	}
	return bitFlags == 1<<26-1
}

func IsPangramUnicode(input string) bool {
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
