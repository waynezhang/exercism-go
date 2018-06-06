package etl

import "strings"

// Transform doc here
func Transform(input map[int][]string) map[string]int {
	ret := make(map[string]int)

	for k, v := range input {
		for _, ch := range v {
			ret[strings.ToLower(ch)] = k
		}
	}

	return ret
}
