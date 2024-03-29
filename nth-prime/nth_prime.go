package prime

import (
	"errors"
	"math"
)

var primes = []int{2}

func isPrimeWithHistory(n int) bool {
	for _, prime := range primes {
		if n%prime == 0 {
			return false
		}
	}
	return true
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n == 1 || n%2 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// NthWithHistory returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func NthWithHistory(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("N must be greater than or equal to 1")
	}
	// Special case for the first prime 2, then we can only test odd numbers starting at 3.
	if n == 1 {
		return 2, nil
	}
	for num := 3; len(primes) < n; num += 2 {
		if isPrimeWithHistory(num) {
			primes = append(primes, num)
		}
	}
	return primes[n-1], nil
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("N must be greater than or equal to 1")
	}
	// Special case for the first prime 2, then we can only test odd numbers starting at 3.
	if n == 1 {
		return 2, nil
	}
	count := 1
	num := 1
	for count < n {
		num += 2
		if isPrime(num) {
			count++
		}
	}
	return num, nil
}
