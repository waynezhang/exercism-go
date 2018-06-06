package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

// AreaCode doc here
func AreaCode(input string) (string, error) {
	str, err := Number(input)
	if err != nil {
		return "", err
	}
	return str[:3], nil
}

// Format doc here
func Format(input string) (string, error) {
	str, err := Number(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", str[:3], str[3:6], str[6:]), nil
}

// Number doc here
func Number(input string) (string, error) {
	str := regexp.MustCompile("[^0-9]*").ReplaceAllString(input, "")

	if len(str) == 11 {
		if str[0] != '1' {
			return "", errors.New("")
		}
		str = str[1:]
	}
	if len(str) != 10 {
		return "", errors.New("")
	}
	if str[0] < '2' || str[3] < '2' {
		return "", errors.New("")
	}

	return str, nil
}
