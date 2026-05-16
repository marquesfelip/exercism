package microblog

func Truncate(phrase string) string {
	runeCount := 0

	for i := range phrase {
		if runeCount == 5 {
			return string([]byte(phrase[:i]))
		}

		runeCount++
	}

	return phrase
}
