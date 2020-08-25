// Package proverb implements… a proverb
package proverb

import "fmt"

// Proverb return a version of the classical proverb adapted to the list of words
func Proverb(rhyme []string) []string {
	verses := []string{}
	rhymeCount := len(rhyme)

	for i := 1; i < rhymeCount; i++ {
		verses = append(verses, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	if rhymeCount != 0 {
		verses = append(verses, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}
	return verses
}
