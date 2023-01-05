package atbash

var cipher = []byte("abcdefghijklmnopqrstuvwxyz")
var cipherMap map[byte]int

func init() {
	cipherMap = make(map[byte]int)
	for idx, c := range cipher {
		cipherMap[c] = idx
	}
}

func Atbash(s string) string {
	cleaned := []byte{}
	for _, c := range []byte(s) {
		switch {
		case c >= '0' && c <= '9':
			cleaned = append(cleaned, c)
		case c >= 'a' && c <= 'z':
			cleaned = append(cleaned, c)
		case c >= 'A' && c <= 'Z':
			cleaned = append(cleaned, c+32)
		}
	}

	encoded := []byte{}
	for idx, c := range cleaned {
		if idx != 0 && idx%5 == 0 {
			encoded = append(encoded, ' ')
		}
		encoded = append(encoded, encode(c))
	}

	return string(encoded)
}

func encode(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return cipher[25-cipherMap[c]]
	}
	return c
}
