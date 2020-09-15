package sorting

// BubbleSort returns a slice of integers sorted in ascending order
// from the input, using the Bubble Sort algorithm. The input is unaltered.
//
// It compares adjacent values 2 by 2 and swaps them if they're not in
// correct order, from j=0 to max=n-1. This leads the highest value to reach
// the last position. The whole slice is sorted by repeating the process
// while decrementing max.
func BubbleSort(input []int) []int {
	n := len(input)
	arr := make([]int, n)
	copy(arr, input)

	for max := n - 1; max > 0; max-- {
		hasSwapped := false

		for j := 0; j < max; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				hasSwapped = true
			}
		}

		// no swap means slice is sorted
		if !hasSwapped {
			break
		}
	}

	return arr
}
