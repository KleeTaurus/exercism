package perfect

import (
	"errors"
	"math"
)

// Define the Classification type here.
type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("Only positive integer is allowed")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return -1, ErrOnlyPositive
	}

	factors := make(map[int64]bool)
	max := math.Ceil(math.Sqrt(float64(n)))
	for i := int64(1); i < int64(max); i++ {
		if n%i == 0 {
			factors[i] = true
			if i != 1 {
				factors[n/i] = true
			}
		}
	}

	var sum int64
	for k := range factors {
		sum += k
	}

	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
