package transpose

func Transpose(input []string) []string {
	maxCols := 0
	rowLength := make([]int, 0)
	for _, row := range input {
		rowLength = append(rowLength, len(row))
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	trans := make([][]byte, maxCols, maxCols)
	for i, row := range input {
		cols := getMaxCol(i, rowLength)

		for j := 0; j < cols; j++ {
			if j < len(row) {
				trans[j] = append(trans[j], row[j])
			} else {
				trans[j] = append(trans[j], byte(' '))
			}
		}
	}

	transStr := make([]string, 0)
	for _, row := range trans {
		transStr = append(transStr, string(row))
	}

	return transStr
}

func getMaxCol(row int, rowLen []int) int {
	max := rowLen[row]
	for i := row + 1; i < len(rowLen); i++ {
		if rowLen[i] > max {
			max = rowLen[i]
		}
	}

	return max
}
