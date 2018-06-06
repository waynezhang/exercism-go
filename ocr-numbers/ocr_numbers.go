package ocr

import (
	"strings"
)

// Recognize doc here
func Recognize(input string) []string {
	lines := strings.Split(input, "\n")[1:]

	ret := make([]string, len(lines)/4)
	for i := 0; i < len(ret); i++ {
		ret[i] = recognizeLine(lines[i*4 : i*4+4])
	}

	return ret
}

func recognizeLine(line []string) string {
	maxWidth := 0
	for _, str := range line {
		if len(str) > maxWidth {
			maxWidth = len(str)
		}
	}

	ret := make([]byte, maxWidth/3)
	for i := range ret {
		if n, ok := recognizeDigit(line, i*3); ok {
			ret[i] = byte('0' + n)
		} else {
			ret[i] = '?'
		}
	}

	return string(ret)
}

func recognizeDigit(input []string, offset int) (int, bool) {
	numbers := []string{
		`
 _
| |
|_|
   `,
		`

  |
  |
   `,
		`
 _
 _|
|_
   `,
		`
 _
 _|
 _|
   `,
		`

|_|
  |
   `,
		`
 _
|_
 _|
   `,
		`
 _
|_
|_|
   `,
		`
 _
  |
  |
   `,
		`
 _
|_|
|_|
   `,
		`
 _
|_|
 _|
   `,
	}

	for val, num := range numbers {
		equal := true
		for i, line := range strings.Split(num, "\n")[1:] {
			src := strings.TrimRight(line, " ")
			dst := strings.TrimRight(input[i][offset:offset+3], " ")
			if src != dst {
				equal = false
				break
			}
		}
		if equal {
			return val, true
		}
	}

	return 0, false
}
