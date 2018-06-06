package atbash

import "strings"

// Atbash doc here
func Atbash(input string) string {
	table := "zyxwvutsrqponmlkjihgfedcba"

	builder := strings.Builder{}

	wrote := 0
	for _, char := range strings.ToLower(input) {
		if char >= '0' && char <= '9' {
		} else if char >= 'a' && char <= 'z' {
			char = rune(table[char-'a'])
		} else {
			continue
		}
		if wrote > 0 && wrote%5 == 0 {
			builder.WriteByte(' ')
		}
		builder.WriteRune(char)
		wrote++
	}

	return builder.String()
}
