package anagram

import (
	"reflect"
	"sort"
	"strings"
)

func sortRunes(word string) []rune {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return runes
}

func Detect(subject string, candidates []string) (anagrams []string) {
	subject = strings.ToLower(subject)
	sortedSubject := sortRunes(subject)
	for _, word := range candidates {
		wordLowercase := strings.ToLower(word)
		if wordLowercase == subject {
			continue
		}
		sortedWord := sortRunes(wordLowercase)
		if reflect.DeepEqual(sortedSubject, sortedWord) {
			anagrams = append(anagrams, word)
		}
	}
	return
}
