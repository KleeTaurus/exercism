package queenattack

import (
	"errors"
	"math"
)

var columns = map[byte]int{
	'a': 0,
	'b': 1,
	'c': 2,
	'd': 3,
	'e': 4,
	'f': 5,
	'g': 6,
	'h': 7,
}

var rows = map[byte]int{
	'8': 0,
	'7': 1,
	'6': 2,
	'5': 3,
	'4': 4,
	'3': 5,
	'2': 6,
	'1': 7,
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if !validate(whitePosition, blackPosition) {
		return false, errors.New("Invalid position")
	}

	var whitePoint [2]int
	var blackPoint [2]int

	whitePoint[0] = columns[whitePosition[0]]
	whitePoint[1] = rows[whitePosition[1]]

	blackPoint[0] = columns[blackPosition[0]]
	blackPoint[1] = rows[blackPosition[1]]

	if whitePoint[0] == blackPoint[0] ||
		whitePoint[1] == blackPoint[1] ||
		math.Abs(float64(whitePoint[0]-blackPoint[0])) == math.Abs(float64(whitePoint[1]-blackPoint[1])) {
		return true, nil
	}

	return false, nil
}

func validate(whitePosition, blackPosition string) bool {
	if len(whitePosition) != 2 || len(blackPosition) != 2 {
		return false
	}

	if whitePosition == blackPosition {
		return false
	}

	if !(whitePosition[0] >= 'a' && whitePosition[0] <= 'h') {
		return false
	}

	if !(blackPosition[0] >= 'a' && blackPosition[0] <= 'h') {
		return false
	}

	if !(whitePosition[1] >= '1' && whitePosition[1] <= '8') {
		return false
	}

	if !(blackPosition[1] >= '1' && blackPosition[1] <= '8') {
		return false
	}

	return true
}
