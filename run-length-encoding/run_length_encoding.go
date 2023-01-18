package encode

import (
	"fmt"
	"math"
)

func RunLengthEncode(input string) string {
	output := make([]byte, 0)
	var current byte
	var amount int
	for i := 0; i < len(input); i++ {
		if i == 0 {
			current = input[i]
			amount = 1
		} else {
			if input[i] != current {
				output = append(output, []byte(encodeLetter(amount, current))...)
				current = input[i]
				amount = 1
			} else {
				amount++
			}
		}

		if i == len(input)-1 {
			output = append(output, []byte(encodeLetter(amount, current))...)
		}

	}

	return string(output)
}

func encodeLetter(amount int, letter byte) string {
	if amount == 1 {
		return fmt.Sprintf("%c", letter)
	}
	return fmt.Sprintf("%d%c", amount, letter)
}

// 12WB12W3B24WB
func RunLengthDecode(input string) string {
	output := make([]byte, 0)
	amount := make([]byte, 0)
	for i := 0; i < len(input); i++ {
		if input[i] >= 48 && input[i] <= 57 {
			amount = append(amount, input[i])
		} else {
			output = append(output, []byte(decodeLetter(amount, input[i]))...)
			amount = make([]byte, 0)
		}
	}

	return string(output)
}

func decodeLetter(amount []byte, letter byte) string {
	if len(amount) == 0 {
		return string(letter)
	}

	decoded := make([]byte, 0)
	num := 0
	for i := 0; i < len(amount); i++ {
		num += int(amount[i]-48) * int(math.Pow10(len(amount)-i-1))
	}

	for i := 0; i < num; i++ {
		decoded = append(decoded, letter)
	}

	return string(decoded)
}
