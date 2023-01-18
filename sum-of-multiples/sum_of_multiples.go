package summultiples

func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]bool)
	for _, d := range divisors {
		for i := 1; ; i++ {
			if d == 0 || d*i >= limit {
				break
			}
			multiples[d*i] = true
		}
	}

	sum := 0
	for k := range multiples {
		sum += k
	}

	return sum
}
