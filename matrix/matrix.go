package matrix

import (
	"bytes"
	"errors"
	"strconv"
)

// Matrix document here
type Matrix [][]int

// Rows document here
func (m *Matrix) Rows() [][]int {
	var ret = make([][]int, len(*m))
	for i, row := range *m {
		ret[i] = make([]int, len(row))
		copy(ret[i], row)
	}
	return ret
}

// Cols document here
func (m *Matrix) Cols() [][]int {
	if len(*m) == 0 {
		return [][]int{}
	}

	var ret = make([][]int, len((*m)[0]))
	for i := range ret {
		ret[i] = make([]int, len(*m))
	}
	for i, row := range *m {
		for j, val := range row {
			ret[j][i] = val
		}
	}
	return ret
}

// Set document here
func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(*m) {
		return false
	}
	if col < 0 || col >= len((*m)[row]) {
		return false
	}
	(*m)[row][col] = val
	return true
}

// New document here
func New(input string) (*Matrix, error) {
	var lines = bytes.Split([]byte(input), []byte("\n"))
	var m = make(Matrix, len(lines))

	for i, line := range lines {
		var chs = bytes.Split(bytes.TrimSpace(line), []byte(" "))
		if i > 0 && len(chs) != len(m[i-1]) {
			return nil, errors.New("")
		}

		m[i] = make([]int, len(chs))
		for j, ch := range chs {
			val, err := strconv.ParseInt(string(ch), 10, 64)
			if err != nil {
				return nil, err
			}
			m[i][j] = int(val)
		}
	}
	return &m, nil
}
