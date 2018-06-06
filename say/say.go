package say

import "strings"

func sayHundreds(n int) string {
	var mapping = map[int]string{
		0:   "",
		1:   "one",
		2:   "two",
		3:   "three",
		4:   "four",
		5:   "five",
		6:   "six",
		7:   "seven",
		8:   "eight",
		9:   "nine",
		10:  "ten",
		11:  "eleven",
		12:  "twelve",
		13:  "thirteen",
		14:  "fourteen",
		15:  "fifteen",
		16:  "sixteen",
		17:  "seventeen",
		18:  "eighteen",
		19:  "nineteen",
		20:  "twenty",
		30:  "thirty",
		40:  "forty",
		50:  "fifty",
		60:  "sixty",
		70:  "seventy",
		80:  "eighty",
		90:  "ninty",
		100: "hundred",
	}

	ret := []string{}
	if n >= 100 {
		ret = []string{mapping[n/100], mapping[100]}
		n %= 100
	}
	if str, ok := mapping[n]; ok {
		ret = append(ret, str)
	} else {
		str := strings.Join([]string{mapping[n/10*10], mapping[n%10]}, "-")
		ret = append(ret, str)
	}

	return strings.Join(ret, " ")
}

// Say doc here
func Say(n int64) (string, bool) {
	if n < 0 || n > 999999999999 {
		return "", false
	}

	if n == 0 {
		return "zero", true
	}

	var tails = []string{
		"", "thousand", "million", "billion", "trillion", "quadrillion", "quintillion",
	}

	var ret = make([]string, 0)
	for _, tail := range tails {
		if n == 0 {
			break
		}
		reminder := int(n % 1000)
		str := sayHundreds(reminder)
		if len(str) > 0 {
			ret = append([]string{str, tail}, ret...)
		}
		n /= 1000
	}

	return strings.TrimSpace(strings.Join(ret, " ")), true
}
