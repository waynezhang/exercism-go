package secret

// Handshake doc here
func Handshake(secret uint) []string {
	var strs = []string{
		"wink",
		"double blink",
		"close your eyes",
		"jump",
	}

	var ret []string
	var reverse = secret&(1<<4) > 0
	secret &= 1<<0 | 1<<1 | 1<<2 | 1<<3
	for i, s := range strs {
		var mask uint = 1 << uint(i)
		if secret&mask > 0 {
			ret = append(ret, s)
		}
	}

	if reverse {
		for i := 0; i < len(ret)/2; i++ {
			ret[i], ret[len(ret)-i-1] = ret[len(ret)-i-1], ret[i]
		}
	}

	return ret
}
