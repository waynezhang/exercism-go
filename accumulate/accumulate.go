package accumulate

// Accumulate doc here
func Accumulate(collection []string, converter func(string) string) []string {
	var result = make([]string, 0, len(collection))

	for _, s := range collection {
		result = append(result, converter(s))
	}

	return result
}
