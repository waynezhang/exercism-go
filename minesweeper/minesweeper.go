package minesweeper

import (
	"errors"
	"regexp"
)

// Count doc here
func (board *Board) Count() error {
	size := len((*board)[0])
	for i, line := range *board {
		var reg *regexp.Regexp
		if i == 0 || i == len(*board)-1 {
			reg = regexp.MustCompile("^\\+-*\\+$")
		} else {
			reg = regexp.MustCompile("^\\|[ *]*\\|$")
		}
		if !reg.Match(line) || len(line) != size {
			return errors.New("")
		}
	}

	for i := 1; i < len(*board)-1; i++ {
		for j := 1; j < size-1; j++ {
			ch := (*board)[i][j]
			if ch == '*' || ch == '|' {
				continue
			}
			count := 0
			for x := i - 1; x <= i+1; x++ {
				for y := j - 1; y <= j+1; y++ {
					if (*board)[x][y] == '*' {
						count++
					}
				}
			}
			if count > 0 {
				(*board)[i][j] = byte('0' + count)
			}
		}
	}
	return nil
}
