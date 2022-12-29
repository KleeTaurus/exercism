package twelve

import "fmt"

var (
	day = []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
		"sixth",
		"seventh",
		"eighth",
		"ninth",
		"tenth",
		"eleventh",
		"twelfth",
	}

	message = []string{
		"a Partridge in a Pear Tree",
		"two Turtle Doves",
		"three French Hens",
		"four Calling Birds",
		"five Gold Rings",
		"six Geese-a-Laying",
		"seven Swans-a-Swimming",
		"eight Maids-a-Milking",
		"nine Ladies Dancing",
		"ten Lords-a-Leaping",
		"eleven Pipers Piping",
		"twelve Drummers Drumming",
	}
)

func Verse(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s",
		day[i-1], dailyMessage(i-1))
}

func Song() string {
	var song string
	for i := 1; i <= len(day); i++ {
		song += Verse(i)
		if i != len(day) {
			song += "\n"
		}
	}
	return song
}

func dailyMessage(i int) string {
	var song string

	if i == 0 {
		song = message[i]
	} else {
		for j := i; j >= 0; j-- {
			song += message[j]
			if j > 1 {
				song += ", "
			} else if j == 1 {
				song += ", and "
			}
		}
	}

	return song + "."
}
