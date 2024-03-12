package isogram

import "unicode"

/*
goos: darwin
goarch: amd64
pkg: isogram
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkIsIsogram
BenchmarkIsIsogram-12            	 3170043	       362.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsIsogramBitfield
BenchmarkIsIsogramBitfield-12    	 3318597	       357.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsIsogram2
BenchmarkIsIsogram2-12           	 2259518	       466.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsIsogramInt
BenchmarkIsIsogramInt-12         	 3042606	       412.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkIsIsogramMap
BenchmarkIsIsogramMap-12         	  226922	      4939 ns/op	    1400 B/op	      13 allocs/op
*/

func IsIsogram(word string) bool {
	found := [26]bool{}
	for _, r := range word {
		letter := unicode.ToLower(r) - 'a'
		if letter >= 0 && letter <= 26 {
			if found[letter] {
				return false
			} else {
				found[letter] = true
			}
		}
	}
	return true
}

func IsIsogramBitfield(word string) bool {
	var found uint32 = 0
	for _, r := range word {
		letter := unicode.ToLower(r) - 'a'
		if letter >= 0 && letter <= 26 {
			if found&(1<<letter) != 0 {
				return false
			} else {
				found |= 1 << letter
			}
		}
	}
	return true
}

func IsIsogram2(word string) bool {
	found := [26]bool{}
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		letter := unicode.ToLower(r) - 'a'
		if found[letter] {
			return false
		} else {
			found[letter] = true
		}
	}
	return true
}

func IsIsogramInt(word string) bool {
	found := [26]int{}
	for _, r := range word {
		letter := unicode.ToLower(r) - 'a'
		if letter >= 0 && letter <= 26 {
			if found[letter] == 1 {
				return false
			} else {
				found[letter] = 1
			}
		}
	}
	return true
}

func IsIsogramMap(word string) bool {
	found := map[rune]bool{}
	for _, r := range word {
		letter := unicode.ToLower(r)
		if letter >= 'a' && letter <= 'z' {
			if _, exists := found[letter]; exists {
				return false
			}
			found[letter] = true
		}
	}
	return true
}
