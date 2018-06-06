package transpose

import (
// "strings"
)

// Transpose doc here
func Transpose(strs []string) []string {
	maxLen := 0

	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	ret := make([]string, maxLen)
	for i := 0; i < maxLen; i++ {
		bytes := make([]byte, len(strs))

		space := -1
		for j, str := range strs {
			if i < len(str) {
				space = -1
				bytes[j] = str[i]
			} else {
				if space < 0 {
					space = j
				}
				bytes[j] = ' '
			}
		}

		if space == -1 {
			ret[i] = string(bytes)
		} else {
			ret[i] = string(bytes[:space])
		}
	}

	return ret
}
