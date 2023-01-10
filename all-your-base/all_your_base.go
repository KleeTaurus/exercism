package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return []int{}, errors.New("input base must be >= 2")
	}

	if outputBase < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}

	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	var number int
	for i := 0; i < len(inputDigits); i++ {
		number += inputDigits[i] * int(math.Pow(float64(inputBase), float64(len(inputDigits)-i-1)))
	}

	output := []int{}
	for {
		rem := number % outputBase
		output = append([]int{rem}, output...)

		number /= outputBase
		if number == 0 {
			break
		}
	}
	return output, nil
}
