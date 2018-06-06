package lsproduct

import (
	"errors"
)

func product(bytes []byte) (int, error) {
	p := 1
	for _, b := range bytes {
		n := int(b - '0')
		if n < 0 || n > 9 {
			return 0, errors.New("")
		}
		p *= n
	}
	return p, nil
}

// LargestSeriesProduct doc here
func LargestSeriesProduct(s string, n int) (int, error) {
	if n < 0 || n > len(s) {
		return 0, errors.New("")
	}

	bytes := []byte(s)
	max := 0
	for i := 0; i < len(s)-n+1; i++ {
		p, err := product(bytes[i : i+n])
		if err != nil {
			return 0, err
		}
		if p > max {
			max = p
		}
	}

	return max, nil
}
