package romannumerals

import (
	"fmt"
	"strings"
)

const (
	minNumber = 1
	maxNumber = 4000
)

// romanNumerals represents a slice of struct that contains the Roman numeral values and their corresponding positions.
// Each struct in the slice has two fields: 'position' indicating the position value of the Roman numeral and 'values'
// indicating the possible strings associated with that position.
var romanNumerals = []struct {
	position int
	values   []string
}{
	{1000, []string{"", "M", "MM", "MMM"}},
	{100, []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}},
	{10, []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}},
	{1, []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}},
}

// ToRomanNumeral takes an integer number as input and returns its Roman numeral representation.
// If the input number is outside the valid range [minNumber, maxNumber), an error is returned.
// minNumber refers to the minimum number that can be represented as a Roman numeral (const minNumber = 1).
// maxNumber refers to the maximum number that can be represented as a Roman numeral (const maxNumber = 4000).
// Roman numeral mappings are stored in the romanNumerals slice.
// The function iterates over the romanNumerals slice, dividing the input number by the current range position
// and appending the corresponding Roman numeral values to the output string.
// The resulting Roman numeral representation and no error is returned if successful.
// If the input number is invalid, an error message is returned.
// Example usage: ToRomanNumeralMapping(24) returns "XXIV", nil
func ToRomanNumeral(number int) (string, error) {
	var out string
	if number < minNumber || number >= maxNumber {
		return "", fmt.Errorf("cannot represent %d as Roman numeral", number)
	}
	for i := 0; i < len(romanNumerals); i++ {
		rangeNumerals := romanNumerals[i]
		count := number / rangeNumerals.position
		number = number % rangeNumerals.position
		out = out + rangeNumerals.values[count]
	}
	return out, nil
}

// numeralToRomanMapping is a slice of struct that represents the mapping between numeral values and their corresponding Roman numerals.
// Each struct in the slice has two fields: 'numeral' indicating the numeral value and 'roman' indicating its corresponding Roman numeral string.
var numeralToRomanMapping = []struct {
	numeral int
	roman   string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

// ToRomanNumeralMapping takes an integer number as input and returns its Roman numeral representation.
// If the input number is outside the valid range [minNumber, maxNumber), an error is returned.
// minNumber refers to the minimum number that can be represented as a Roman numeral.
// maxNumber refers to the maximum number that can be represented as a Roman numeral.
// numeralToRomanMapping is a slice that stores the Roman numeral mappings as structs with numeral and roman fields.
// The function iterates over the numeralToRomanMapping slice, dividing the input number by the current numeral value
// and appending the corresponding Roman numeral strings to the output string using strings.Repeat.
// The resulting Roman numeral representation and no error is returned if successful.
// If the input number is invalid, an error message is returned.
// Example usage: ToRomanNumeralMapping(24) returns "XXIV", nil
func ToRomanNumeralMapping(number int) (string, error) {
	var out string
	if number < minNumber || number >= maxNumber {
		return "", fmt.Errorf("cannot represent %d as Roman numeral", number)
	}
	for i := 0; i < len(numeralToRomanMapping); i++ {
		count := number / numeralToRomanMapping[i].numeral
		out += strings.Repeat(numeralToRomanMapping[i].roman, count)
		number -= count * numeralToRomanMapping[i].numeral
	}
	return out, nil
}
