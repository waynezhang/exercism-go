package armstrong

import (
	"fmt"
	"math"
)

// IsNumber doc here
func IsNumber(num int) bool {
	str := fmt.Sprintf("%d", num)
	power := float64(len(str))
	sum := 0
	for _, d := range str {
		sum += int(math.Pow(float64(d-'0'), power))
		if sum > num {
			return false
		}
	}

	return sum == num
}
