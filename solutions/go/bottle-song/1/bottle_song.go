package bottlesong

import (
	"fmt"
	"strings"
)

var bottleMap = map[int]string{
	0:  "No",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

func Recite(startBottles, takeDown int) []string {
	var song = make([]string, 0, (4*takeDown + (takeDown - 1)))

	for i := startBottles; i > startBottles-takeDown; i-- {
		song = append(song, verse(i)...)

		if i-1 > startBottles-takeDown {
			song = append(song, "")
		}
	}

	return song
}

func verse(bottle int) []string {
	var verses = make([]string, 0)

	curBottle := bottleMap[bottle]
	remBottle := bottleMap[bottle-1]
	bottleHelper := pluralize(bottle)

	verses = append(verses, fmt.Sprintf("%s green %s hanging on the wall,", curBottle, bottleHelper))
	verses = append(verses, fmt.Sprintf("%s green %s hanging on the wall,", curBottle, bottleHelper))
	verses = append(verses, "And if one green bottle should accidentally fall,")

	switch bottle - 1 {
	case 0:
		verses = append(verses, "There'll be no green bottles hanging on the wall.")
	default:
		verses = append(verses, fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(remBottle), pluralize(bottle-1)))
	}

	return verses
}

func pluralize(num int) string {
	if num == 1 {
		return "bottle"
	}

	return "bottles"
}
