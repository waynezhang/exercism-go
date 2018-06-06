package sublist

// Relation document here
type Relation string

func contains(list1, list2 []int) bool {
	for i := range list2 {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// Sublist document here
func Sublist(list1, list2 []int) Relation {
	var flipped = false
	if len(list1) < len(list2) {
		list1, list2 = list2, list1
		flipped = true
	}

	var contained = false
	for i := 0; i < len(list1)-len(list2)+1; i++ {
		if contains(list1[i:], list2) {
			contained = true
			break
		}
	}

	if contained {
		if len(list1) == len(list2) {
			return "equal"
		}
		if flipped {
			return "sublist"
		}
		return "superlist"
	}
	return "unequal"
}
