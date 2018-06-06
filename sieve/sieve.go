package sieve

func generate(output chan int) {
	for i := 2; ; i++ {
		output <- i
	}
}

func filter(input chan int, output chan int, prime int) {
	for {
		n := <-input
		if n%prime != 0 {
			output <- n
		}
	}
}

// Sieve doc here
func Sieve(n int) []int {
	numbers := make(chan int)
	go generate(numbers)

	ret := make([]int, 0)
	for {
		prime := <-numbers
		if prime > n {
			break
		}
		ret = append(ret, prime)
		filtered := make(chan int)
		go filter(numbers, filtered, prime)
		numbers = filtered
	}

	return ret
}
