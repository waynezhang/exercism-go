package stringset

import (
	"fmt"
	"strings"
)

// Set doc here
type Set struct {
	elements map[string]bool
}

// String doc here
func (s Set) String() string {
	strs := make([]string, 0, len(s.elements))
	for str := range s.elements {
		strs = append(strs, fmt.Sprintf("\"%s\"", str))
	}
	return fmt.Sprintf("{%s}", strings.Join(strs, ", "))
}

// IsEmpty doc here
func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

// Has doc here
func (s Set) Has(ele string) bool {
	return s.elements[ele]
}

// Subset doc here
func Subset(s1, s2 Set) bool {
	return len(Difference(s1, s2).elements) == 0
}

// Disjoint doc here
func Disjoint(s1, s2 Set) bool {
	return len(Difference(s1, s2).elements) == len(s1.elements)
}

// Equal doc here
func Equal(s1, s2 Set) bool {
	intersections := len(Intersection(s1, s2).elements)
	return intersections == len(s1.elements) && intersections == len(s2.elements)
}

// Add doc here
func (s Set) Add(ele string) {
	s.elements[ele] = true
}

// Intersection doc here
func Intersection(s1, s2 Set) Set {
	s := New()
	for ele := range s1.elements {
		if s2.Has(ele) {
			s.Add(ele)
		}
	}
	return s
}

// Difference doc here
func Difference(s1, s2 Set) Set {
	s := New()
	for ele := range s1.elements {
		if !s2.Has(ele) {
			s.Add(ele)
		}
	}
	return s
}

// Union doc here
func Union(s1, s2 Set) Set {
	s := New()
	for ele := range s1.elements {
		s.Add(ele)
	}
	for ele := range s2.elements {
		s.Add(ele)
	}
	return s
}

// NewFromSlice doc here
func NewFromSlice(slice []string) Set {
	s := Set{}
	s.elements = make(map[string]bool)
	for _, str := range slice {
		s.Add(str)
	}
	return s
}

// New doc here
func New() Set {
	return NewFromSlice(nil)
}
