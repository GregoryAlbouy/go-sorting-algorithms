package sorting

// CountingSort returns a slice of integers sorted in ascending order
// from the input, using the Counting Sort algorithm. The input is unaltered.
//
// This version is slighlty changed to accept negative numbers. If the input
// contains negative values, |min| is added to the size of the counting slice
// and the negative values are counted on indexes following the positive values
// (i.e. i > max).
// example:
//  input:         [1, -3, 3, 1, -2]
//  values:         0(x0), 1(x2), 2(x0), 3(x1), -1(x0), -2(x0), -3(x1)
//  counts slice:  [0, 2, 0, 1,  0,  1,  1]
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

	for i := 0; i < n; i++ {
		v := input[i]
		if v < 0 {
			// positivize v (indexes can't be negative) and add max as offset
			counts[max+(-v)]++
		} else {
			counts[v]++
		}
	}

	// append negative integers to final array
	// fetching the negative section of counts in reverse order
	// because higher is lesser when values are negated {7,42} => {-42,-7}
	for i := countsLen - 1; i > max; i-- {
		count := counts[i]

		if count == 0 {
			continue
		}

		// remove the offset (-max) and re-negate the value
		v := -(i - max)

		for k := 0; k < count; k++ {
			arr = append(arr, v)
		}
	}

	// append positive integers to final array
	for v, count := range counts[:max+1] {
		if count == 0 {
			continue
		}

		for i := 0; i < count; i++ {
			arr = append(arr, v)
		}
	}

	return arr
}
