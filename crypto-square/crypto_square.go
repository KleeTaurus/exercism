package cryptosquare

import (
	"math"
)

func Encode(pt string) string {
	nt := normalize(pt)
	cols, rows := calcColsAndRows(nt)
	return encrypt(nt, cols, rows)
}

func normalize(pt string) []byte {
	nt := []byte{}

	for _, c := range []byte(pt) {
		switch {
		case c >= '0' && c <= '9':
			nt = append(nt, c)
		case c >= 'a' && c <= 'z':
			nt = append(nt, c)
		case c >= 'A' && c <= 'Z':
			nt = append(nt, c+32)
		}
	}

	return nt
}

func calcColsAndRows(input []byte) (int, int) {
	sr := math.Sqrt(float64(len(input)))
	if sr-float64(int(sr)) > 0.5 {
		return int(math.Ceil(sr)), int(math.Ceil(sr))
	} else {
		return int(math.Ceil(sr)), int(math.Floor(sr))
	}
}

func encrypt(input []byte, cols, rows int) string {
	secret := []byte{}
	for i := 0; i < cols; i++ {
		for j := 0; j < rows; j++ {
			if j*cols+i < len(input) {
				secret = append(secret, input[j*cols+i])
			} else {
				secret = append(secret, ' ')
			}
		}
		if i != cols-1 {
			secret = append(secret, ' ')
		}
	}

	return string(secret)
}
