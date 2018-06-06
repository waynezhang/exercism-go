package pythagorean

import "math"

// Triplet doc here
type Triplet []int

// Range doc here
func Range(min, max int) []Triplet {
	ret := make([]Triplet, 0)

	for i := min; i < max-2; i++ {
		for j := i + 1; j < max-1; j++ {
			n := math.Sqrt(float64(i*i + j*j))
			if n == math.Ceil(n) {
				ret = append(ret, Triplet{i, j, int(n)})
			}
		}
	}
	return ret
}

// Sum doc here
func Sum(p int) []Triplet {
	triplets := make([]Triplet, 0)
	for _, t := range Range(1, p) {
		if t[0]+t[1]+t[2] == p {
			triplets = append(triplets, t)
		}
	}
	return triplets
}
