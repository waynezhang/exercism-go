package twofer

func ShareWith(s string) string {
	if len(s) == 0 {
		s = "you"
	}
	return "One for " + s + ", one for me."
}
