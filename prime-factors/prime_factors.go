package prime

func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	var i int64

	for i = 2; i <= n; i++ {
		for {
			if n%i != 0 {
				break
			}
			n /= i
			factors = append(factors, i)

			if n == 1 {
				return factors
			}
		}
	}

	return []int64{}
}
