package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, fmt.Errorf("n should be at least 1")
	}
	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
