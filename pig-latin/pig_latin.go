package piglatin

import (
	"strings"
)

// Sentence doc here
func Sentence(sentence string) string {
	converter := func(word string) string {
		qu := strings.Index(word, "qu")
		y := strings.Index(word, "y")
		vowel := strings.IndexAny(word, "aeiou")
		if len(word) >= 2 && (word[:2] == "xr" || word[:2] == "yt") {
			vowel = 0
		}

		index := 0

		if vowel == 0 {
			index = 0
		} else if y == 1 && len(word) == 2 {
			word = strings.Join([]string{word[1:], word[:1]}, "")
			index = 0
		} else if y > 0 && (y < vowel || vowel == -1) {
			index = y
		} else if qu >= 0 {
			index = qu + 2
		} else {
			index = vowel
		}
		return strings.Join([]string{word[index:], word[:index], "ay"}, "")
	}

	words := strings.Split(sentence, " ")
	converted := make([]string, len(words))
	for i, word := range words {
		converted[i] = converter(word)
	}

	return strings.Join(converted, " ")
}
