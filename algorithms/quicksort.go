package algorithms

import "sync"

// QuickSort returns a slice of integers sorted in ascending order
// from the input, using the Quick Sort algorithm. The input is unaltered.
func QuickSort(input []int) []int {
	return runQuickSort(input, false)
}

// QuickSortConc returns a slice of integers sorted in ascending order
// from the input, using the Quick Sort algorithm. The input is unaltered.
//
// Unlike the regular Quick Sort algorithm, it uses concurrency to accelerate
// the process (~ +20% performance).
func QuickSortConc(input []int) []int {
	return runQuickSort(input, true)
}

func runQuickSort(input []int, isConc bool) []int {
	n := len(input)
	arr := make([]int, n)
	copy(arr, input)

	return quickSort(arr, 0, n-1, isConc)
}

func quickSort(arr []int, left, right int, isConc bool) []int {
	if left < right {
		pivotIndex := pivot(arr, left, right)

		if !isConc {
			quickSort(arr, left, pivotIndex-1, false)
			quickSort(arr, pivotIndex+1, right, false)
		} else {
			wg := sync.WaitGroup{}
			wg.Add(2)
			go func() { quickSort(arr, left, pivotIndex-1, false); wg.Done() }()
			go func() { quickSort(arr, pivotIndex+1, right, false); wg.Done() }()
			wg.Wait()
		}
	}

	return arr
}

func pivot(arr []int, start, end int) int {
	pivot := arr[start]
	swapIndex := start

	for i := start + 1; i <= end; i++ {
		if pivot > arr[i] {
			swapIndex++
			swap(arr, i, swapIndex)
		}
	}

	swap(arr, start, swapIndex)

	return swapIndex
}
