package reverse

func Reverse(input string) string {
	letters := []rune(input)
	reversedArray := []rune{}

	for i := len(letters) - 1; i >= 0; i-- {
		reversedArray = append(reversedArray, letters[i])
	}

	return string(reversedArray)
}
