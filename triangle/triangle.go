package triangle

import (
	"math"
	"sort"
)

// Kind doc here
type Kind int

const (
	// NaT not a triangle
	NaT = iota
	// Equ equilateral
	Equ
	// Iso isosceles
	Iso
	// Sca scalene
	Sca
)

func isSide(n float64) bool {
	return !math.IsNaN(n) && !math.IsInf(n, 0) && n > 0
}

// KindFromSides doc here
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	var sides = []float64{a, b, c}
	sort.Sort(sort.Float64Slice(sides))

	if sides[0]+sides[1] < sides[2] {
		return NaT
	}
	if !isSide(sides[0]) || !isSide(sides[2]) {
		return NaT
	}

	if a == b && a == c {
		k = Equ
	} else if a == b || b == c || a == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
