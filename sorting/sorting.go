package sorting

type sortFunc func([]int) []int

// Algorithm represents a sorting algorithm. It has a name and a sorting
// function.
type Algorithm struct {
	Name     string
	sortFunc sortFunc
}

// Sorter implements a function Sort([]int) []int
type Sorter interface {
	Sort([]int) []int
}

// Sort returns a copy of input slice sorted in ascending order.
func (a Algorithm) Sort(arr []int) []int {
	return a.sortFunc(arr)
}

// AllAlgorithms is a slice referencing all sorting algorithms in this package.
var AllAlgorithms = []Algorithm{
	{"BubbleSort", BubbleSort},
	{"SelectionSort", SelectionSort},
	{"InsertionSort", InsertionSort},
	{"MergeSort", MergeSort},
	{"MergeSortConc", MergeSortConc},
	{"QuickSort", QuickSort},
	{"QuickSortConc", QuickSortConc},
	{"RadixSort", RadixSort},
	{"GoSort", GoSort},
	{"CountingSort", CountingSort},
}
