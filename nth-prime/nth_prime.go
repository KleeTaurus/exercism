package prime

import (
	"errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n can't less than 1")
	}

	founded := 0
	for i := 2; ; i++ {
		if isPrime(i) {
			founded++
			if founded == n {
				return i, nil
			}
		}
	}
}

func isPrime(n int) bool {
	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
