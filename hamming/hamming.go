package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("")
	}

	count := 0
	for i := 0; i < len(a); i++ {
		if a[i:i+1] != b[i:i+1] {
			count++
		}
	}

	return count, nil
}
