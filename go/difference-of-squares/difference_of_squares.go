package diffsquares

import (
	"math"
)

func SquareOfSum(n int) int {
	var sum int = 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(n int) int {
	var sum int = 0

	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}

	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
