package reverse

import (
	"unicode/utf8"
)

func String(str string) string {
	var len = utf8.RuneCountInString(str)
	var buffer = make([]rune, len)
	for _, ch := range str {
		len--
		buffer[len] = ch
	}
	return string(buffer)
}
