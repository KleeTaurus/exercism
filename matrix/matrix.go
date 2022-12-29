package matrix

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix = *matrix

type matrix struct {
	values [][]int
	rows   int
	cols   int
}

var (
	rowRe = regexp.MustCompile("\n")
	colRe = regexp.MustCompile("\\s+")
)

func New(s string) (Matrix, error) {
	m := &matrix{}

	rows := rowRe.Split(s, -1)
	m.rows = len(rows)
	m.values = make([][]int, m.rows)

	for i := 0; i < m.rows; i++ {
		cols := colRe.Split(strings.TrimSpace(rows[i]), -1)

		if i == 0 {
			m.cols = len(cols)
		} else {
			if m.cols != len(cols) {
				return nil, fmt.Errorf("Wrong number of columns, wanted: %d, got: %d", m.cols, len(cols))
			}
		}

		m.values[i] = make([]int, m.cols)
		for j := 0; j < m.cols; j++ {
			v, err := strconv.Atoi(cols[j])
			if err != nil {
				return nil, err
			}
			m.values[i][j] = v
		}
	}

	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cpy := make([][]int, m.cols)
	for i := 0; i < m.cols; i++ {
		cpy[i] = make([]int, m.rows)
		for j := 0; j < m.rows; j++ {
			cpy[i][j] = m.values[j][i]
		}
	}
	return cpy
}

func (m Matrix) Rows() [][]int {
	cpy := make([][]int, m.rows)
	for i := 0; i < m.rows; i++ {
		cpy[i] = make([]int, m.cols)
		for j := 0; j < m.cols; j++ {
			cpy[i][j] = m.values[i][j]
		}
	}
	return cpy
}

func (m Matrix) Set(row, col, val int) bool {
	if row >= 0 && row < m.rows && col >= 0 && col < m.cols {
		m.values[row][col] = val
		return true
	}

	return false
}
