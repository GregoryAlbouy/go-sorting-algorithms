package sorting

type sortFunc func([]int) []int

type sortFuncMap map[string]sortFunc

// Algorithm represents a sorting algorithm. It has a name and a sorting
// function.
type Algorithm struct {
	Name string
	Sort sortFunc
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
