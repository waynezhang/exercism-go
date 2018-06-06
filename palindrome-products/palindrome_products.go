package palindrome

import (
	"errors"
	"fmt"
)

// Product doc here
type Product struct {
	Product        int
	Factorizations [][2]int
}

func isPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Products doc here
func Products(fmin, fmax int) (pmin Product, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}

	for i := fmin; i < fmax+1; i++ {
		for j := i; j < fmax+1; j++ {
			p := i * j
			if !isPalindrome(p) {
				continue
			}
			if p < pmin.Product || len(pmin.Factorizations) == 0 {
				pmin.Product = p
				pmin.Factorizations = [][2]int{[2]int{i, j}}
			} else if p == pmin.Product {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
			}
			if p > pmax.Product || len(pmax.Factorizations) == 0 {
				pmax.Product = p
				pmax.Factorizations = [][2]int{[2]int{i, j}}
			} else if p == pmax.Product {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
			}
		}
	}

	if len(pmin.Factorizations) == 0 {
		err = errors.New("no palindromes")
	}

	return
}
