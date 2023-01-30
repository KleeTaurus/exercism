package wordsearch

import "errors"

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	for _, word := range words {
		for i := 0; i < len(puzzle); i++ {
			for j := 0; j < len(puzzle[i]); j++ {
				if puzzle[i][j] == word[0] {
					matched, x, y := ifWordMatched(word, puzzle, i, j)
					if matched {
						result[word] = [2][2]int{{j, i}, {y, x}}
					}
				}
			}
		}
	}

	if len(result) != len(words) {
		return nil, errors.New("Not all words were founded")
	}

	return result, nil
}

func ifWordMatched(word string, puzzle []string, i, j int) (bool, int, int) {
	// find for 8 directions
	wordLen := len(word)
	if j+wordLen <= len(puzzle[i]) {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i][j+x]
		}
		if string(bw) == word {
			return true, i, j + wordLen - 1
		}
	}
	if j-wordLen+1 >= 0 {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i][j-x]
		}
		if string(bw) == word {
			return true, i, j - wordLen + 1
		}
	}
	if i+wordLen <= len(puzzle) {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i+x][j]
		}
		if string(bw) == word {
			return true, i + wordLen - 1, j
		}
	}
	if i-wordLen+1 >= 0 {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i-x][j]
		}
		if string(bw) == word {
			return true, i - wordLen + 1, j
		}
	}
	if j+wordLen <= len(puzzle[i]) && i+wordLen <= len(puzzle) {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i+x][j+x]
		}
		if string(bw) == word {
			return true, i + wordLen - 1, j + wordLen - 1
		}
	}
	if j-wordLen+1 >= 0 && i-wordLen+1 >= 0 {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i-x][j-x]
		}
		if string(bw) == word {
			return true, i - wordLen + 1, j - wordLen + 1
		}
	}
	if j+wordLen <= len(puzzle[i]) && i-wordLen+1 >= 0 {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i-x][j+x]
		}
		if string(bw) == word {
			return true, i - wordLen + 1, j + wordLen - 1
		}
	}
	if j-wordLen+1 >= 0 && i+wordLen <= len(puzzle) {
		bw := make([]byte, wordLen)
		for x := 0; x < wordLen; x++ {
			bw[x] = puzzle[i+x][j-x]
		}
		if string(bw) == word {
			return true, i + wordLen - 1, j - wordLen + 1
		}
	}
	return false, -1, -1
}
