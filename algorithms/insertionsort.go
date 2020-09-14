package algorithms

// InsertionSort returns a slice of integers sorted in ascending order
// from the input, using the Bubble Sort algorithm. The input is unaltered.
func InsertionSort(input []int) []int {
	n := len(input)
	arr := make([]int, n)
	copy(arr, input)

	for i := 1; i < n; i++ {
		cur := arr[i]
		j := i - 1

		for ; j >= 0 && arr[j] > cur; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = cur
	}

	return arr
}
