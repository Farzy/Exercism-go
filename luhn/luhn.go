package luhn

import "unicode"

func Valid(id string) bool {
	digitCount := 0
	sum := 0
	doubleDigit := false

	runes := []rune(id)
	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if unicode.IsNumber(r) {
			digit := r - '0'
			if doubleDigit {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += int(digit)
			digitCount++
			doubleDigit = !doubleDigit
		} else if r != ' ' {
			return false
		}
	}
	if digitCount <= 1 || (sum%10) != 0 {
		return false
	}
	return true
}
