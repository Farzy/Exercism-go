package diffsquares

func SquareOfSum(n int) int {
	if n < 0 {
		panic("n must be a positive number")
	}
	return n * n * (n + 1) * (n + 1) / 4
}

func SumOfSquares(n int) int {
	if n < 0 {
		panic("n must be a positive number")
	}
	return n * (n + 1) * (2*n + 1) / 6
}

func SumOfSquaresSimple(n int) (sum int) {
	if n < 0 {
		panic("n must be a positive number")
	}
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func DifferenceSimple(n int) int {
	return SquareOfSum(n) - SumOfSquaresSimple(n)
}
