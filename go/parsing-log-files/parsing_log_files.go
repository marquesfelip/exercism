package parsinglogfiles

import (
	"regexp"
)

func IsValidLine(text string) bool {
	validInitLine := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)

	return validInitLine.MatchString(text)
}

func SplitLogLine(text string) []string {
	return regexp.MustCompile(`<(\~*|\**|\=*|\-*)*?>`).Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int = 0

	for _, line := range lines {
		if regexp.MustCompile(`(?mi)".*(password).*"`).MatchString(line) {
			count++
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	expression := regexp.MustCompile(`(?m)(.*)(User)(\s+)(\w+)(.*)`)

	for index, line := range lines {
		if expression.MatchString(line) {
			lines[index] = expression.ReplaceAllString(line, `[USR] $4 $1$2$3$4$5`)
		}
	}

	return lines
}
