package cryptosquare

import (
	"bytes"
	"math"
	"regexp"
	"strings"
)

// Encode doc here
func Encode(s string) string {
	s = strings.ToLower(s)
	reg := regexp.MustCompile("[^0-9a-z]*")
	s = reg.ReplaceAllString(s, "")

	if len(s) == 0 {
		return ""
	}

	c := int(math.Ceil(math.Sqrt(float64(len(s)))))
	r := len(s) / c
	if len(s)%c > 0 {
		r++
	}

	chars := []byte(s)
	buffer := bytes.Buffer{}
	for i := 0; i < c; i++ {
		if i > 0 {
			buffer.WriteByte(' ')
		}
		for j := 0; j < r; j++ {
			idx := i + j*c
			if idx < len(chars) {
				buffer.WriteByte(chars[idx])
			} else {
				buffer.WriteByte(' ')
			}
		}
	}

	return buffer.String()
}
