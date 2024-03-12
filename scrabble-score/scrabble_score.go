package scrabble

// goos: darwin
// goarch: amd64
// pkg: scrabble
// cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
// BenchmarkScore
// BenchmarkScore-12                 	 3931598	       290.0 ns/op	       0 B/op	       0 allocs/op
// BenchmarkScoreToUpperString
// BenchmarkScoreToUpperString-12    	 1848860	       609.9 ns/op	      80 B/op	       9 allocs/op
// BenchmarkScoreMap
// BenchmarkScoreMap-12              	  110658	     10423 ns/op	    6315 B/op	      50 allocs/op
// BenchmarkScoreSwitch
// BenchmarkScoreSwitch-12           	 5403943	       214.4 ns/op	       0 B/op	       0 allocs/op

import (
	"strings"
	"unicode"
)

// B': 3,
// C': 3,
// A': 1,
// D': 2,
// E': 1,
// F': 4,
// G': 2,
// H': 4,
// I': 1,
// J': 8,
// K': 5,
// L': 1,
// M': 3,
// N': 1,
// O': 1,
// P': 3,
// Q': 10,
// R': 1,
// S': 1,
// T': 1,
// U': 1,
// V': 4,
// W': 4,
// X': 8,
// Y': 4,
// Z': 10,

func Score(word string) (score int) {
	scores := []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

	for _, letter := range word {
		score += scores[unicode.ToUpper(letter)-'A']
	}
	return
}

func ScoreToUpperString(word string) (score int) {
	scores := []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

	for _, letter := range strings.ToUpper(word) {
		score += scores[letter-'A']
	}
	return
}

func ScoreMap(word string) (score int) {
	scores := map[rune]int{
		'A': 1,
		'B': 3,
		'C': 3,
		'D': 2,
		'E': 1,
		'F': 4,
		'G': 2,
		'H': 4,
		'I': 1,
		'J': 8,
		'K': 5,
		'L': 1,
		'M': 3,
		'N': 1,
		'O': 1,
		'P': 3,
		'Q': 10,
		'R': 1,
		'S': 1,
		'T': 1,
		'U': 1,
		'V': 4,
		'W': 4,
		'X': 8,
		'Y': 4,
		'Z': 10,
	}

	for _, letter := range strings.ToUpper(word) {
		score += scores[letter]
	}
	return
}

func ScoreSwitch(word string) (score int) {
	for _, letter := range word {
		switch unicode.ToLower(letter) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}
	return score
}
