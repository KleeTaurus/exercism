package sieve

func Sieve(limit int) []int {
	m := make(map[int]bool)
	primes := make([]int, 0)

	for i := 2; i <= limit; i++ {
		if m[i] == false {
			primes = append(primes, i)
			for j := 2; j*i <= limit; j++ {
				m[i*j] = true
			}
		}
	}

	return primes
}
