package wordcount

import (
	"regexp"
	"strings"
	"unicode"
)

// My solution is much faster than the version using Regex:
// goos: darwin
// goarch: amd64
// pkg: wordcount
// cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
// BenchmarkWordCount
// BenchmarkWordCount-12         	  136604	      8671 ns/op	    4032 B/op	      93 allocs/op
// BenchmarkWordCountRegex
// BenchmarkWordCountRegex-12    	   22996	     54123 ns/op	   45214 B/op	     455 allocs/op

// Frequency represents a mapping of string words to their integer frequencies.
type Frequency map[string]int

// isLetter returns true if the given character is a letter, a number, or an apostrophe (if inWord is true).
func isLetter(chr rune, inWord bool) bool {
	return unicode.IsLetter(chr) || unicode.IsNumber(chr) || (inWord && chr == '\'')
}

// makeWord trims the trailing apostrophe from the given strings.Builder and returns a string.
func makeWord(sb strings.Builder) string {
	return strings.TrimSuffix(sb.String(), "'")
}

// WordCount takes a string `phrase` and returns a map
// of word frequencies in that phrase.
//
// It converts the phrase to lowercase using `strings.ToLower`.
//
// It initializes an empty map `frequencies` of type `Frequency`
// (which is a map[string]int).
//
// It uses a boolean variable `inWord` to keep track if it is currently
// inside a word or not.
//
// It uses a `strings.Builder` named `currentWord` to build the current word
// as it iterates over the characters in the phrase.
//
// It iterates over each character `chr` in the phrase.
// If the character is a letter (according to the `isLetter` function),
// it sets `inWord` to true, appends the character `chr` to `currentWord`.
// If the character is not a letter, it checks if it was previously inside a word.
// If it was, it sets `inWord` to false, adds 1 to the frequency of the word
// represented by the contents of `currentWord` in the `frequencies` map,
// resets `currentWord` to an empty string.
//
// After iterating over all characters, it handles the last word by checking if
// `currentWord` is not empty and adds 1 to its frequency in the `frequencies`
// map.
//
// Finally, it returns the `frequencies` map.
//
// Example usage:
// frequencies := WordCount("one fish two fish red fish blue fish")
// fmt.Println(frequencies) // Output: map[blue:1 fish:4 one:1 red:1 two:1]
//
// Note: This function relies on the `makeWord` and `isLetter` functions which are not
// provided in this documentation.
func WordCount(phrase string) Frequency {
	phrase = strings.ToLower(phrase)
	frequencies := make(Frequency)

	inWord := false
	currentWord := strings.Builder{}

	for _, chr := range phrase {
		if isLetter(chr, inWord) {
			inWord = true
			currentWord.WriteRune(chr)
		} else if inWord {
			inWord = false
			frequencies[makeWord(currentWord)] += 1
			currentWord.Reset()
		}
	}
	// Last word
	if currentWord.Len() != 0 {
		frequencies[makeWord(currentWord)] += 1
	}

	return frequencies
}

func WordCountRegex(phrase string) Frequency {
	re := regexp.MustCompile(`\w+('\w+)?`)
	f := make(Frequency)
	for _, s := range re.FindAllString(strings.ToLower(phrase), -1) {
		f[s] += 1
	}
	return f
}
