package prime

// Nth doc here
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}

	primes := []int{2}
	for i := 3; len(primes) < n; i += 2 {
		isPrime := true
		for _, p := range primes {
			if i%p == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes[n-1], true
}
