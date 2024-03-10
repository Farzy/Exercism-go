// Package proverb implements… a proverb
package proverb

// Proverb return a version of the classical proverb adapted to the list of words
func Proverb(rhyme []string) []string {
	var verses []string
	rhymeCount := len(rhyme)

	for i := 1; i < rhymeCount; i++ {
		verses = append(verses, "For want of a "+rhyme[i-1]+" the "+rhyme[i]+" was lost.")
	}
	if rhymeCount != 0 {
		verses = append(verses, "And all for the want of a "+rhyme[0]+".")
	}
	return verses
}
