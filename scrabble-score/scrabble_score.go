package scrabble

import "strings"

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

	for _, letter := range strings.ToUpper(word) {
		score += scores[letter-'A']
	}
	return
}
