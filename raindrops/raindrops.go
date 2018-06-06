package raindrops

import (
	"bytes"
	"fmt"
)

func Convert(rain int) string {
	var buffer bytes.Buffer
	if rain%3 == 0 {
		buffer.WriteString("Pling")
	}
	if rain%5 == 0 {
		buffer.WriteString("Plang")
	}
	if rain%7 == 0 {
		buffer.WriteString("Plong")
	}
	if buffer.Len() == 0 {
		buffer.WriteString(fmt.Sprintf("%d", rain))
	}
	return buffer.String()
}
