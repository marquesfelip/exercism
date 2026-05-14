package armstrongnumbers

func IsNumber(n int) bool {
	var numDigits = make([]int, 0)
	original := n
	result := 0

	for n > 0 {
		numDigits = append(numDigits, (n % 10))
		n /= 10
	}

	for _, digit := range numDigits {
		result += calcPow(digit, len(numDigits))
	}

	return result == original
}

func calcPow(base, power int) int {
	// for i := 0; i < power; i++ {
	result := base
	for range power - 1 {
		result = result * base
	}

	return result
}
