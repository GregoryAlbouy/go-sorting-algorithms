package algorithms

// BubbleSort returns a slice of integers sorted in ascending order
// from the input, using the Bubble Sort algorithm. The input is unaltered.
func BubbleSort(input []int) []int {
	n := len(input)
	arr := make([]int, n)
	copy(arr, input)

	for i := n - 1; i > 0; i-- {
		hasSwapped := false

		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
				hasSwapped = true
			}
		}

		if !hasSwapped {
			break
		}
	}

	return arr
}
