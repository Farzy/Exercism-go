package romannumerals

import (
	"fmt"
)

const (
	minNumber = 1
	maxNumber = 4000
)

var romanNumerals = []struct {
	position int
	values   []string
}{
	{1000, []string{"", "M", "MM", "MMM"}},
	{100, []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}},
	{10, []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}},
	{1, []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}},
}

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
