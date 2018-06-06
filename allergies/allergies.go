package allergies

// AllergicTo doc here
func AllergicTo(score uint, food string) bool {
	for _, f := range Allergies(score) {
		if f == food {
			return true
		}
	}
	return false
}

// Allergies doc here
func Allergies(score uint) []string {
	all := []string{
		"eggs",
		"peanuts",
		"shellfish",
		"strawberries",
		"tomatoes",
		"chocolate",
		"pollen",
		"cats",
	}

	allergies := make([]string, 0)
	for i, food := range all {
		if score&(uint(1)<<uint(i)) > 0 {
			allergies = append(allergies, food)
		}
	}
	return allergies
}
