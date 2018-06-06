package series

// All doc here
func All(n int, s string) []string {
	var bytes = []byte(s)
	var ret = make([]string, len(bytes)-n+1)
	for i := 0; i < len(ret); i++ {
		ret[i] = string(bytes[i : i+n])
	}
	return ret
}

// UnsafeFirst doc here
func UnsafeFirst(n int, s string) string {
	r, _ := First(n, s)
	return r
}

// First doc here
func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return string([]byte(s)[:n]), true
}
