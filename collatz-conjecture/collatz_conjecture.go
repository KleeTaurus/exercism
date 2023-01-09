package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be a positive number")
	}
	steps := 0
	for {
		if n == 1 {
			break
		}
		n = doMath(n)
		steps++
	}

	return steps, nil
}

func doMath(n int) int {
	if n%2 == 0 {
		return n >> 1
	}

	return n<<1 + n + 1
}
