package letter

func ConcurrentFrequency(array []string) FreqMap {
	var results = make(chan FreqMap, len(array))
	for _, str := range array {
		go func(str string) {
			results <- Frequency(str)
		}(str)
	}

	var ret = FreqMap{}
	for range array {
		m := <-results
		for k, v := range m {
			ret[k] += v
		}
	}
	return ret
}
