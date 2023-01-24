package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("Invalid char")
	}

	n := int(char - 'A')
	diamond := make([]string, 0)
	for i := 0; i < 2*n+1; i++ {
		row := emptyByteSlice(2*n + 1)
		if i <= n {
			row[n-i] = 'A' + byte(i)
			row[n+i] = 'A' + byte(i)
		} else {
			row[n-(2*n-i)] = 'A' + byte(2*n-i)
			row[n+(2*n-i)] = 'A' + byte(2*n-i)
		}
		diamond = append(diamond, string(row))
	}

	return strings.Join(diamond, "\n"), nil
}

func emptyByteSlice(n int) []byte {
	bs := make([]byte, n, n)
	for i := 0; i < n; i++ {
		bs[i] = ' '
	}
	return bs
}
