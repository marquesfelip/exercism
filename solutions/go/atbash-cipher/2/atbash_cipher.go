package atbashcipher

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	count := 0
	var b strings.Builder
	const atbashRate = 'a' + 'z'

	for _, r := range s {
		r = unicode.ToLower(r)

		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			if count > 0 && count%5 == 0 {
				b.WriteRune(' ')
			}

			if r >= 'a' && r <= 'z' {
				b.WriteRune(atbashRate - r)
			} else {
				b.WriteRune(r)
			}

			count++
		}
	}

	return b.String()
}
