package brackets

// Bracket doc here
func Bracket(s string) (bool, error) {
	bracketMap := map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
	brackets := make([]rune, 0)
	for _, r := range s {
		switch r {
		case '[', '{', '(':
			brackets = append(brackets, r)
		case ']', '}', ')':
			if len(brackets) > 0 && brackets[len(brackets)-1] == bracketMap[r] {
				brackets = brackets[:len(brackets)-1]
			} else {
				brackets = append(brackets, r)
				break
			}
		}
	}
	return len(brackets) == 0, nil
}
