package atbash

const (
	offset = 97
)

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
		return byte(25 - (int(c) - offset) + offset)
	}
	return c
}
