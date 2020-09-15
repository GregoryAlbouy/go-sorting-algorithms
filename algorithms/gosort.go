package algorithms

import "sort"

// GoSort returns a slice of integers sorted in ascending order
// from the input, using the Go builtin sort algorithm.
// The input is unaltered.
func GoSort(input []int) []int {
	n := len(input)
	arr := make([]int, n)
	copy(arr, input)

	sort.Ints(arr)

	return arr
}
