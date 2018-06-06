package wordcount

import (
	"strings"
	"unicode"
)

// Frequency doc here
type Frequency map[string]int

// WordCount doc here
func WordCount(s string) Frequency {
	frequency := Frequency{}

	splitted := strings.FieldsFunc(strings.ToLower(s), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
	})
	for _, word := range splitted {
		word := strings.Trim(word, "'")
		frequency[word]++
	}
	return frequency
}
