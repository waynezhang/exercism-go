package prime

// Factors doc here
func Factors(num int64) []int64 {
	factors := []int64{}

	for num%2 == 0 {
		factors = append(factors, 2)
		num /= 2
	}
	for i := int64(3); i*i <= num; i += 2 {
		for num%i == 0 {
			factors = append(factors, i)
			num /= i
		}
	}
	if num > 1 {
		factors = append(factors, num)
	}

	return factors
}
