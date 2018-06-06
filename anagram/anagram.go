package anagram

import (
	"sort"
	"strings"
)

func resort(s string) string {
	bytes := []byte(strings.ToLower(s))
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}

// Detect doc here
func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	resorted := resort(subject)

	ret := make([]string, 0)
	for _, s := range candidates {
		if resort(s) == resorted && subject != strings.ToLower(s) {
			ret = append(ret, s)
		}
	}
	return ret
}
