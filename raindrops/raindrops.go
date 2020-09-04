package raindrops

import "strconv"

// Convert converts a number into a string that contains raindrop sounds
//
// Note: A version using strings.Builder is about 33% slower and uses 3 times more
// memory allocations.
func Convert(num int) (result string) {
	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		result = strconv.Itoa(num)
	}

	return
}
