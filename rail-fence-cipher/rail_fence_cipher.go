package railfence

func makeIndexes(length int, railNum int) []int {
	rails := make([][]int, railNum)

	curr := 0
	direction := 1
	for i := 0; i < length; i++ {
		rails[curr] = append(rails[curr], i)
		curr += direction
		if curr == 0 || curr == railNum-1 {
			direction *= -1
		}
	}

	ret := make([]int, 0, length)
	for _, rail := range rails {
		ret = append(ret, rail...)
	}
	return ret
}

// Encode doc here
func Encode(s string, n int) string {
	indexes := makeIndexes(len(s), n)
	str := make([]byte, len(s))
	for i, idx := range indexes {
		str[i] = s[idx]
	}
	return string(str)
}

// Decode doc here
func Decode(s string, n int) string {
	indexes := makeIndexes(len(s), n)
	str := make([]byte, len(s))
	for i, idx := range indexes {
		str[idx] = s[i]
	}
	return string(str)
}
