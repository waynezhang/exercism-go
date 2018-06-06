package binarysearch

// SearchInts doc here
func SearchInts(array []int, val int) int {
	if len(array) == 0 {
		return -1
	}

	start, end, mid := 0, len(array)-1, (len(array)-1)/2
	for {
		mid = (start + end) / 2
		if mid == start {
			if array[start] == val {
				return start
			} else if array[end] == val {
				return end
			} else {
				return -1
			}
		}
		if array[mid] < val {
			start = mid
		} else {
			end = mid
		}
	}
}
