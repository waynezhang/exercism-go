package luhn

import "strings"

func Valid(str string) bool {
	str = strings.Replace(str, " ", "", -1)
	if len(str) < 2 {
		return false
	}

	var sum = 0
	for i, ch := range str {
		val := int(ch - '0')
		if (len(str)-i)%2 == 0 {
			product := val * 2
			if product > 9 {
				product -= 9
			}
			sum += product
		} else {
			sum += val
		}
	}

	return sum%10 == 0
}
