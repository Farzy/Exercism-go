package grains

import "errors"

func Square(number int) (uint64, error) {
	if number > 0 && number <= 64 {
		return 1 << (number - 1), nil
	} else {
		return 0, errors.New("number must be between 1 and 64")
	}
}

func Total() (sum uint64) {
	for i := 1; i <= 64; i++ {
		sq, _ := Square(i)
		sum += sq
	}
	return
}
