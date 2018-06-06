package isogram

import "strings"

func IsIsogram(str string) bool {
	var maps = map[rune]int{}

	for _, ch := range strings.ToLower(str) {
		maps[ch] += 1
		if maps[ch] > 1 && strings.ContainsRune("abcdefghijklmnopqrstuvwxyz", ch) {
			return false
		}
	}

	return true
}
