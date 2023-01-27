package bottlesong

import (
	"fmt"
	"strings"
)

var numbers = []string{"no", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

var template = []string{
	"%s green %s hanging on the wall,",
	"And if one green bottle should accidentally fall,",
	"There'll be %s green %s hanging on the wall.",
}

func Recite(startBottles, takeDown int) []string {
	lyrics := make([]string, 0)
	for i := 0; i < takeDown; i++ {
		number := numbers[startBottles-i]
		bottles := []string{"bottles", "bottles"}
		if number == "Two" {
			bottles[1] = "bottle"
		} else if number == "One" {
			bottles[0] = "bottle"
		}

		lyrics = append(lyrics, fmt.Sprintf(template[0], number, bottles[0]))
		lyrics = append(lyrics, fmt.Sprintf(template[0], number, bottles[0]))
		lyrics = append(lyrics, template[1])
		lyrics = append(lyrics, fmt.Sprintf(template[2], strings.ToLower(numbers[startBottles-i-1]), bottles[1]))

		if i != takeDown-1 {
			lyrics = append(lyrics, "")
		}
	}

	return lyrics
}
