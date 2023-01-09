package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
	input := n
	digits := make([]int, 0)

	for {
		rem := n % 10
		n = n / 10

		digits = append(digits, rem)
		if n == 0 {
			break
		}
	}

	nums := len(digits)

	var sum float64
	for _, digit := range digits {
		sum += math.Pow(float64(digit), float64(nums))
	}

	return input == int(sum)
}
