package isbn

func IsValidISBN(isbn string) bool {
	digits := []int{}
	for _, c := range []byte(isbn) {
		switch {
		case c >= '0' && c <= '9':
			digits = append(digits, int(c-48))
		case c == 'X' && len(digits) == 9:
			digits = append(digits, 10)
		case c == '-':
			continue
		default:
			return false
		}
	}

	if len(digits) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < len(digits); i++ {
		sum += digits[i] * (10 - i)
	}

	return sum%11 == 0
}
