package strain

// Ints doc here
type Ints []int

// Lists doc here
type Lists [][]int

// Strings doc here
type Strings []string

// Keep doc here
func (ints Ints) Keep(predict func(int) bool) Ints {
	if ints == nil {
		return nil
	}

	newInts := Ints{}
	for _, v := range ints {
		if predict(v) {
			newInts = append(newInts, v)
		}
	}

	return newInts
}

// Discard doc here
func (ints Ints) Discard(predict func(int) bool) Ints {
	if ints == nil {
		return nil
	}

	newInts := Ints{}
	for _, v := range ints {
		if !predict(v) {
			newInts = append(newInts, v)
		}
	}

	return newInts
}

// Keep doc here
func (lists Lists) Keep(predict func([]int) bool) Lists {
	if lists == nil {
		return nil
	}

	newLists := Lists{}
	for _, l := range lists {
		if predict(l) {
			newLists = append(newLists, l)
		}
	}

	return newLists
}

// Keep doc here
func (strings Strings) Keep(predict func(string) bool) Strings {
	if strings == nil {
		return nil
	}

	newStrings := Strings{}
	for _, s := range strings {
		if predict(s) {
			newStrings = append(newStrings, s)
		}
	}

	return newStrings
}
