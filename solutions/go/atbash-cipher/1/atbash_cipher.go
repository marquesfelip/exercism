package atbashcipher

import "strings"

func Atbash(s string) string {
	count := 0
	var b strings.Builder

	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			if count > 0 && count%5 == 0 {
				b.WriteRune(' ')
			}

			if r >= 'a' && r <= 'z' {
				b.WriteRune(('a' + 'z') - r)
			} else if r >= '0' && r <= '9' {
				b.WriteRune(r)
			}

			count++
		}
	}

	return b.String()
}
