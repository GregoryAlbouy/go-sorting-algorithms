package sorting

// CountingSort returns a slice of integers sorted in ascending order
// from the input, using the Counting Sort algorithm. The input is unaltered.
//
// This version is slighlty changed to accept negative numbers. If the input
// contains negative values, |min| is added to the size of the counting slice
// and the negative values are counted on indexes following the positive values
// (i.e. i > max).
// e.g.: INPUT [1, -3, 3, 1, -2] => COUNT [0, 2, 0, 1,  0,  1,  1]
//                                         ^  ^  ^  ^   ^   ^   ^
//                                  VALUES 0, 1, 2, 3, -1, -2, -3
func CountingSort(input []int) []int {
	n := len(input)
	if n == 0 {
		return input
	}

	var arr []int
	min, max := minmax(input)
	countsLen := max + 1

	// if necessary, add space to store negative values
	if min < 0 {
		countsLen += -min
	}

	counts := make([]int, countsLen)
	posCounts := counts[:max+1]
	// negCounts := counts[max+1:]

	for i := 0; i < n; i++ {
		v := input[i]
		if v < 0 {
			// index the absolute value with offset = max
			counts[max+(-v)]++
		} else {
			counts[v]++
		}
	}

	for i := countsLen - 1; i > max; i-- {
		count := counts[i]

		if count == 0 {
			continue
		}

		v := -(i - max)

		for j := 0; j < count; j++ {
			arr = append(arr, v)
		}
	}

	for v, count := range posCounts {
		if count == 0 {
			continue
		}

		for i := 0; i < count; i++ {
			arr = append(arr, v)
		}
	}

	return arr
}
