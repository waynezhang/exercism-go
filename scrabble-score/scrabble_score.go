package scrabble

import "strings"

func Score(letters string) int {
	var rawVals = map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}

	var vals = map[rune]int{}
	for k, v := range rawVals {
		for _, ch := range k {
			vals[ch] = v
		}
	}

	var ret = 0
	for _, ch := range strings.ToUpper(letters) {
		ret += vals[ch]
	}

	return ret
}
