package summultiples

// SumMultiples doc here
func SumMultiples(limit int, divisors ...int) int {
	ret := 0

	for i := 1; i < limit; i++ {
		for _, n := range divisors {
			if i%n == 0 {
				ret += i
				break
			}
		}
	}
	return ret
}
