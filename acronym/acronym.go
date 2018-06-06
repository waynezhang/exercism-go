package acronym

import (
	"bytes"
	"strings"
)

func Abbreviate(s string) string {
	array := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})

	var buffer bytes.Buffer
	for _, word := range array {
		buffer.WriteString(word[:1])
	}
	return strings.ToUpper(buffer.String())
}
