package romannumerals

import (
	"errors"
	"strings"
)

// ToRomanNumeral doc here
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", errors.New("")
	}

	chars := [][]byte{
		[]byte("XVIII"),
		[]byte("CLXXX"),
		[]byte("MDCCC"),
		[]byte("  MMM"),
	}
	ret := []string{}
	for _, symbols := range chars {
		reminder := n % 10
		val := ""
		switch reminder {
		case 9:
			val = string(append(symbols[2:3], symbols[0]))
		case 5, 6, 7, 8:
			val = string(append(symbols[1:2], symbols[2:2+reminder-5]...))
		case 4:
			val = string(append(symbols[2:3], symbols[1]))
		case 0, 1, 2, 3:
			val = string(symbols[2 : 2+reminder])
		}
		ret = append([]string{val}, ret...)
		n /= 10
	}
	return strings.Join(ret, ""), nil
}
