package rotationalcipher

var (
	alphabetLower = []byte("abcdefghijklmnopqrstuvwxyz")
	alphabetUpper = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	alphabetMap = make(map[byte]int)
)

func init() {
	for i := 0; i < 26; i++ {
		alphabetMap[alphabetLower[i]] = i
		alphabetMap[alphabetUpper[i]] = i
	}
}

func RotationalCipher(plain string, shiftKey int) string {
	cipher := []byte{}
	for _, c := range []byte(plain) {
		switch {
		case c >= 'a' && c <= 'z':
			cipher = append(cipher, alphabetLower[shift(c, shiftKey)])
		case c >= 'A' && c <= 'Z':
			cipher = append(cipher, alphabetUpper[shift(c, shiftKey)])
		default:
			cipher = append(cipher, c)
		}
	}

	return string(cipher)
}

func shift(c byte, shiftKey int) int {
	idx := alphabetMap[c]
	return (idx + shiftKey) % 26
}
