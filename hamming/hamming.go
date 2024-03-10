// Package hamming help compute the Hamming distance between 2 DNA strands
package hamming

import "fmt"

// Distance returns the Hamming distance betweek 2 DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("string a (%d) & b (%d) are not of equal length", len(a), len(b))
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
