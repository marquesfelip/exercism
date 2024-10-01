package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid input")
	}

	steps := 0

	for n != 1 {
		if n % 2 == 0 {
			n = n / 2
			steps++
		} else {
			n = 3 * n + 1
			steps++
		}
	}

	return steps, nil
}
