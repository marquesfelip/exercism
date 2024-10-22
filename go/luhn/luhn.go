package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	if len(strings.ReplaceAll(id, " ", "")) <= 1 {
		return false
	}

	digits := strings.ReplaceAll(id, " ", "")
	sum := 0
	double := false

	for _, ch := range digits {
		if !unicode.IsDigit(ch) {
			return false
		}
	}

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i] - '0'

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += int(digit)
		double = !double
	}

	return sum % 10 == 0
}
