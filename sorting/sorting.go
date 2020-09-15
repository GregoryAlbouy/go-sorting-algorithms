package sorting

type sortFunc func([]int) []int

// Algorithm represents a sorting algorithm. It has a name and a sorting
// function.
type Algorithm struct {
	Name string
	Sort sortFunc
}

// Sorter implements a function Sort([]int) []int
type Sorter interface {
	Sort([]int) []int
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
}
