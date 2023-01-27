package matrix

import (
	"sort"
	"strconv"
	"strings"
)

// Define the Matrix and Pair types here.
type Matrix struct {
	rows int
	cols int
	m    [][]int
}

type Pair struct {
	row int // start from 1
	col int // start from 1
}

func New(s string) (*Matrix, error) {
	matrix := &Matrix{}

	if s == "" {
		return matrix, nil
	}

	rows := strings.Split(s, "\n")
	matrix.rows = len(rows)
	matrix.m = make([][]int, len(rows))

	for ridx, row := range rows {
		cols := strings.Split(row, " ")
		matrix.cols = len(cols)
		matrix.m[ridx] = make([]int, len(cols))
		for cidx, col := range cols {
			v, _ := strconv.Atoi(col)
			matrix.m[ridx][cidx] = v
		}
	}
	// fmt.Printf("matrix: %v\n", matrix)
	return matrix, nil
}

func (m *Matrix) Saddle() []Pair {
	ps := []Pair{}
	for i := 0; i < m.rows; i++ {
		cpy := make([]int, m.cols)
		copy(cpy, m.m[i])
		sort.Ints(cpy)
		rowmax := cpy[m.cols-1]
		// fmt.Println("row: ", m.m[i], ", max: ", rowmax)

		for j := 0; j < m.cols; j++ {
			if m.m[i][j] == rowmax {
				colmin := rowmax
				for k := 0; k < m.rows; k++ {
					if k == i {
						continue
					}

					if m.m[k][j] < rowmax {
						colmin = m.m[k][j]
						break
					}
				}
				if colmin == rowmax {
					ps = append(ps, Pair{i + 1, j + 1})
				}
			}
		}
	}
	return ps
}
