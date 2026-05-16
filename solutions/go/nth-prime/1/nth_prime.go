package nthprime

import (
	"errors"
	"math"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("there is no zeroth prime")
	}

	candidate := 2
	count := 1

	for count < n {
		candidate++

		if isPrime(candidate) {
			count++
		}
	}

	return candidate, nil
}

func isPrime(candidate int) bool {
	if candidate < 2 {
		return false
	}

	if candidate == 2 {
		return true
	}

	if candidate%2 == 0 {
		return false
	}

	limit := int(math.Sqrt(float64(candidate)))

	for i := 3; i <= limit; i += 2 {
		if candidate%i == 0 {
			return false
		}
	}

	return true
}
