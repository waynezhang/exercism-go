package pascal

// Triangle doc here
func Triangle(size int) [][]int {
	var ret = make([][]int, size, size)

	for i := 0; i < size; i++ {
		ret[i] = make([]int, i+1, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 {
				ret[i][j] = 1
			} else if j < i {
				ret[i][j] = ret[i-1][j-1] + ret[i-1][j]
			} else {
				ret[i][j] = ret[i-1][j-1]
			}
		}
	}

	return ret
}
