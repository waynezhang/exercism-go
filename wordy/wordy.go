package wordy

import (
	"regexp"
	"strconv"
	"strings"
)

// Answer doc here
func Answer(question string) (result int, ok bool) {
	question = regexp.MustCompile("What is (.+)\\?").ReplaceAllString(question, "${1}")
	tokens := strings.Split(question, " ")

	if val, err := strconv.Atoi(tokens[0]); err == nil {
		result = int(val)
	} else {
		return 0, false
	}

	addFunc := func(a, b int) int { return a + b }
	minusFunc := func(a, b int) int { return a - b }
	multipleFunc := func(a, b int) int { return a * b }
	divideFunc := func(a, b int) int { return a / b }

	for i := 1; i < len(tokens); {
		opFunc := addFunc
		switch tokens[i] {
		case "plus":
			opFunc = addFunc
			i++
		case "minus":
			opFunc = minusFunc
			i++
		case "multiplied":
			opFunc = multipleFunc
			i += 2
		case "divided":
			opFunc = divideFunc
			i += 2
		}

		if val, err := strconv.Atoi(tokens[i]); err == nil {
			i++
			result = opFunc(result, int(val))
		} else {
			return 0, false
		}
	}

	return result, true
}
