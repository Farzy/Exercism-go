package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := make(map[string]int, 26)
	for value, letters := range in {
		for _, letter := range letters {
			out[strings.ToLower(letter)] = value
		}
	}
	return out
}
