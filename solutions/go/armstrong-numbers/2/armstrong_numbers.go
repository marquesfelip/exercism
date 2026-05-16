package armstrongnumbers

func IsNumber(n int) bool {
	original := n
	digits := countDigits(n)

	sum := 0

	for n > 0 {
		d := n % 10
		sum += calcPow(d, digits)
		n /= 10
	}

	return sum == original
}

func countDigits(n int) int {
	count := 0
	for n > 0 {
		n /= 10
		count++
	}

	return count
}

func calcPow(base, power int) int {
	result := base
	for range power - 1 {
		result = result * base
	}

	return result
}
