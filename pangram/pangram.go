package pangram

func IsPangram(input string) bool {
	m := make(map[byte]bool)

	letters := []byte(input)
	for _, l := range letters {
		switch {
		case l >= 65 && l <= 90:
			m[l+32] = true
		case l >= 97 && l <= 122:
			m[l] = true
		}
	}

	return len(m) == 26
}
