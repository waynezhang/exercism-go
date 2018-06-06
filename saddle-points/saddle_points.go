package matrix

import (
	"strconv"
	"strings"
)

// Pair doc here
type Pair struct {
	row int
	col int
}

// Matrix doc here
type Matrix [][]int

// Saddle doc here
func (m *Matrix) Saddle() []Pair {
	pairs := make([]Pair, 0)

	for i, row := range *m {

		maxJ := 0
		for j := range row {
			if row[maxJ] < row[j] {
				maxJ = j
			}
		}
		minI := i
		for r := range *m {
			if (*m)[r][maxJ] < (*m)[minI][maxJ] {
				minI = r
				break
			}
		}
		if minI == i {
			pairs = append(pairs, Pair{row: i, col: maxJ})
		}
	}
	return pairs
}

// New doc here
func New(input string) (*Matrix, error) {
	m := Matrix{}
	for _, row := range strings.Split(input, "\n") {
		cols := make([]int, 0)
		for _, col := range strings.Split(row, " ") {
			if val, err := strconv.Atoi(col); err == nil {
				cols = append(cols, val)
			} else {
				return nil, err
			}
		}
		m = append(m, cols)
	}
	return &m, nil
}
