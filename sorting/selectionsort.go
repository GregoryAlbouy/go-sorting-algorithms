package sorting

// SelectionSort sorts a slice of integers by ascending order
// using the Selection Sort algorithm.
func SelectionSort(input []int) []int {
	n := len(input)
	arr := make([]int, n)
	copy(arr, input)

	for i := 0; i < n-1; i++ {
		vmin := arr[i]
		imin := i
		for j := i; j < n; j++ {
			if arr[j] < vmin {
				vmin = arr[j]
				imin = j
			}
		}
		if i != imin {
			swap(arr, i, imin)
		}
	}

	return arr
}
