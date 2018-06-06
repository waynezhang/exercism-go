package pangram

import "strings"

// IsPangram doc here
func IsPangram(s string) bool {
	s = strings.ToLower(s)

	hash := make(map[rune]bool)
	for _, ch := range s {
		if 'a' <= ch && ch <= 'z' {
			hash[ch] = true
		}
	}

	return len(hash) == 26
}
