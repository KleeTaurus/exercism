package romannumerals

import (
	"errors"
)

var (
	lettersMap = [][]byte{
		[]byte{'I', 'V'},
		[]byte{'X', 'L'},
		[]byte{'C', 'D'},
		[]byte{'M'},
	}
)

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("The maximum number supported is 3,999")
	}
	digits := toDigits(input)

	var romanNumerals string
	for i, d := range digits {
		letters := digitToLetters(d, i)
		romanNumerals = letters + romanNumerals
	}
	return romanNumerals, nil
}

func toDigits(input int) []int {
	digits := []int{}
	for {
		r := input % 10
		digits = append(digits, r)
		if input = input / 10; input == 0 {
			break
		}
	}
	return digits
}

func digitToLetters(digit int, idx int) string {
	var letters string
	switch digit {
	case 1:
		letters = string([]byte{lettersMap[idx][0]})
	case 2:
		letters = string([]byte{lettersMap[idx][0], lettersMap[idx][0]})
	case 3:
		letters = string([]byte{lettersMap[idx][0], lettersMap[idx][0], lettersMap[idx][0]})
	case 4:
		letters = string([]byte{lettersMap[idx][0], lettersMap[idx][1]})
	case 5:
		letters = string([]byte{lettersMap[idx][1]})
	case 6:
		letters = string([]byte{lettersMap[idx][1], lettersMap[idx][0]})
	case 7:
		letters = string([]byte{lettersMap[idx][1], lettersMap[idx][0], lettersMap[idx][0]})
	case 8:
		letters = string([]byte{lettersMap[idx][1], lettersMap[idx][0], lettersMap[idx][0], lettersMap[idx][0]})
	case 9:
		letters = string([]byte{lettersMap[idx][0], lettersMap[idx+1][0]})
	}

	return letters
}
