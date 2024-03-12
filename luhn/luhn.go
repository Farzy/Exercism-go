package luhn

/*
goos: darwin
goarch: amd64
pkg: luhn
cpu: Intel(R) Core(TM) i7-8850H CPU @ 2.60GHz
BenchmarkValid
BenchmarkValid-12            	 3870295	       291.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkValidRune
BenchmarkValidRune-12        	 1553277	       766.0 ns/op	     176 B/op	       1 allocs/op
BenchmarkValidExercism
BenchmarkValidExercism-12    	  501820	      2196 ns/op	    1024 B/op	      36 allocs/op
*/

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) bool {
	var digitCount, sum int
	var doubleDigit bool

	for i := len(id) - 1; i >= 0; i-- {
		c := id[i]
		if c == ' ' {
			continue
		}
		digit := int(c - '0')
		if digit >= 0 && digit <= 9 {
			if doubleDigit {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
			digitCount++
			doubleDigit = !doubleDigit
		} else {
			return false
		}
	}
	return digitCount > 1 && (sum%10) == 0
}

func ValidRune(id string) bool {
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

// This example code from https://exercism.org/tracks/go/exercises/luhn/dig_deeper does
// not even pass the long number test because of an integer overflow!
func ValidExercism(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	if _, err := strconv.Atoi(num); err != nil {
		return false
	}
	total := 0
	pos := 0
	for i := len(num) - 1; i > -1; i-- {
		digit := int(num[i] - '0')
		if pos%2 == 0 {
			total += digit
		} else {
			switch digit {
			case 1, 2, 3, 4:
				total += digit << 1
			case 5, 6, 7, 8:
				total += (digit << 1) - 9
			case 9:
				total += digit
			}
		}
		pos++
	}
	return pos > 1 && total%10 == 0
}
