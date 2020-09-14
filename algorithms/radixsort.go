package algorithms

// RadixSort returns a slice of integers sorted in ascending order
// from the input, using the Radix Sort algorithm. The input is unaltered.
func RadixSort(input []int) []int {
	n := len(input)
	arr := make([]int, n)
	copy(arr, input)

	maxDigitCount := mostDigits(arr)

	for k := 0; k < maxDigitCount; k++ {
		buckets := make([][]int, 10)

		for i := 0; i < n; i++ {
			digit := digitAt(arr[i], k)
			bucket := &buckets[digit]
			*bucket = append(*bucket, arr[i])
		}

		arr = []int{}
		for _, b := range buckets {
			arr = append(arr, b...)
		}
	}

	return arr
}
