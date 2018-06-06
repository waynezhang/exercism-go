package diamond

import (
	"errors"
	"fmt"
	"strings"
)

func padding(length int) string {
	const padding = "                                                    "
	return padding[:length]
}

// Gen doc here
func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", errors.New("")
	}

	halfWidth := int(b - 'A')
	// need add a empty line at the end to pass the test
	ret := make([]string, halfWidth*2+2)

	ret[0] = fmt.Sprintf("%s%c%s", padding(halfWidth), 'A', padding(halfWidth))
	ret[halfWidth*2] = ret[0]

	for i := 1; i <= halfWidth; i++ {
		ret[i] = fmt.Sprintf("%s%c%s%c%s", padding(halfWidth-i), 'A'+i, padding(i*2-1), 'A'+i, padding(halfWidth-i))
		ret[halfWidth*2-i] = ret[i]
	}
	return strings.Join(ret, "\n"), nil
}
