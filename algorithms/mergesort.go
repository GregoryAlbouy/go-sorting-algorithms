package algorithms

import (
	"sync"
)

// MergeSort returns a slice of integers sorted in ascending order
// from the input, using the Merge Sort algorithm. The input is unaltered.
//
// It recursively splits in half the input slice until there's
// n slices of length 1, then the parts are merged back in ascending
// order.
func MergeSort(arr []int) []int {
	return mergeSort(arr, false)
}

// MergeSortConc returns a slice of integers sorted in ascending order
// from the input, using the Merge Sort algorithm. The input is unaltered.
//
// Unlike the regular MergeSort algorithm, it uses concurrency to accelerate
// the process (~ +45% performance).
func MergeSortConc(arr []int) []int {
	return mergeSort(arr, true)
}

func mergeSort(arr []int, isConc bool) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	mid := n / 2
	var left, right []int

	// Regular mode (sequential)

	if !isConc {
		left = mergeSort(arr[mid:], false)
		right = mergeSort(arr[:mid], false)

		return merge(left, right)
	}

	// Concurrent mode

	wg := sync.WaitGroup{}
	wg.Add(2)

	// To avoid the creation of n goroutines which slows down the process,
	// I set a length threshold. Below that limit, concurrency is stopped.
	//
	// It performs 25% better than cancelling concurrency after first call
	// and 400% better than no cancellation at all.
	//
	// The value 1000 is the result of empirical tests.
	const concThreshold = 1000
	if n < concThreshold {
		isConc = false
	}

	go func() {
		left = mergeSort(arr[:mid], isConc)
		wg.Done()
	}()

	go func() {
		right = mergeSort(arr[mid:], isConc)
		wg.Done()
	}()

	wg.Wait()

	return merge(left, right)
}

func merge(a, b []int) []int {
	var arr []int
	m, n := len(a), len(b)
	i, j := 0, 0

	for i < m && j < n {
		if a[i] < b[j] {
			arr = append(arr, a[i])
			i++
		} else {
			arr = append(arr, b[j])
			j++
		}
	}

	for i < m {
		arr = append(arr, a[i])
		i++
	}

	for j < n {
		arr = append(arr, b[j])
		j++
	}

	return arr
}
