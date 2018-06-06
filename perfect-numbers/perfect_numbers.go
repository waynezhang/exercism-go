package perfect

import (
	"errors"
	"math"
)

// ErrOnlyPositive doc here
var ErrOnlyPositive = errors.New("")

// Classification doc here
type Classification int

const (
	// ClassificationPerfect doc here
	ClassificationPerfect = iota
	// ClassificationAbundant doc here
	ClassificationAbundant
	// ClassificationDeficient doc here
	ClassificationDeficient
)

// Classify doc here
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	if n == 1 {
		return ClassificationDeficient, nil
	}

	sum := int64(1)
	for i := int64(2); i <= int64(math.Sqrt(float64(n))) && i < n; i++ {
		if n%i == 0 {
			sum += i
			if i != 1 && n/i != i {
				sum += n / i
			}
		}
	}

	var ret Classification = ClassificationPerfect
	if sum > n {
		ret = ClassificationAbundant
	} else if sum < n {
		ret = ClassificationDeficient
	}
	return ret, nil
}
