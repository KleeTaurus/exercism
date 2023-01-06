package rotationalcipher

const (
	upperOffset = 65
	lowerOffset = 97
)

func RotationalCipher(plain string, shiftKey int) string {
	cipher := []byte{}
	for _, c := range []byte(plain) {
		switch {
		case c >= 'a' && c <= 'z':
			cipher = append(cipher, shift(c, lowerOffset, shiftKey))
		case c >= 'A' && c <= 'Z':
			cipher = append(cipher, shift(c, upperOffset, shiftKey))
		default:
			cipher = append(cipher, c)
		}
	}

	return string(cipher)
}

func shift(c byte, offset int, shiftKey int) byte {
	return byte((int(c)-offset+shiftKey)%26 + offset)
}
