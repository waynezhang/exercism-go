package bob

import "strings"

func Hey(remark string) string {
	answer := "Whatever."

	remark = strings.Trim(remark, "\r\n\t ")
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	lastChar := remark[len(remark)-1:]
	if strings.ToUpper(remark) == remark && strings.ContainsAny(remark, "abcdefghijklmopqrstuvwxyzABCDEFGHIJKLMOPQRSTUVWXYZ") {
		if lastChar == "?" {
			answer = "Calm down, I know what I'm doing!"
		} else {
			answer = "Whoa, chill out!"
		}
	} else if lastChar == "?" {
		answer = "Sure."
	}
	return answer
}
