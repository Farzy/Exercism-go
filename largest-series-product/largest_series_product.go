package lsproduct

import (
	"fmt"
	"unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	numDigits := len(digits)
	if span < 1 || span > numDigits {
		return 0, fmt.Errorf("span of %d is larger than the digits string (%d)", span, numDigits)
	}
	var maxProduct int64
	for i := 0; i <= numDigits-span; i++ {
		product := int64(1)
		for j := i; j < i+span; j++ {
			digitRune := rune(digits[j])
			if !unicode.IsNumber(digitRune) {
				return 0, fmt.Errorf("found non digit in string: '%c'", digitRune)
			}
			digit := int64(digitRune - '0')
			product *= digit
		}
		if product > maxProduct {
			maxProduct = product
		}
	}
	return maxProduct, nil
}
