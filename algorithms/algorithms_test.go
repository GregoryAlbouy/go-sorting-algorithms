package algorithms

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	clog "github.com/gregoryalbouy/go-custom-log"
)

var benchsize = 10000

type testcase struct {
	input       []int
	expected    []int
	description string
}

type testSortingAlgorithm struct {
	name      string
	sortFunc  sortFunc
	testFunc  func(*testing.T)
	benchFunc func(*testing.B)
}

var sortingAlgorithms = []testSortingAlgorithm{
	{"BubbleSort", BubbleSort, TestBubbleSort, BenchmarkBubbleSort},
	{"SelectionSort", SelectionSort, TestSelectionSort, BenchmarkSelectionSort},
	{"InsertionSort", InsertionSort, TestInsertionSort, BenchmarkInsertionSort},
	{"MergeSort", MergeSort, TestMergeSort, BenchmarkMergeSort},
	{"MergeSortConc", MergeSortConc, TestMergeSortConc, BenchmarkMergeSortConc},
	{"QuickSort", QuickSort, TestQuickSort, BenchmarkQuickSort},
	{"QuickSortConc", QuickSortConc, TestQuickSortConc, BenchmarkQuickSortConc},
}

func runTest(t *testing.T, a testSortingAlgorithm) {
	testcases := []testcase{
		{
			description: "positive ints",
			input:       []int{5, 2, 4, 1, 3},
			expected:    []int{1, 2, 3, 4, 5},
		}, {
			description: "relative ints",
			input:       []int{0, 42, -7, -42, 7},
			expected:    []int{-42, -7, 0, 7, 42},
		}, {
			description: "empty slice",
			input:       []int{},
			expected:    []int{},
		}, {
			description: "single value",
			input:       []int{0},
			expected:    []int{0},
		},
	}

	for _, tc := range testcases {
		clone := make([]int, len(tc.input))
		copy(clone, tc.input)

		got := a.sortFunc(tc.input)

		check(t, tc, got, clone)
	}

	if t.Failed() {
		fmt.Println(clog.Red("FAIL"), a.name)
	} else {
		fmt.Println(clog.Green("PASS"), a.name)
	}
}

func runBenchmark(b *testing.B, sfm sortFuncMap, inputLen ...int) {
	const defLen = 100000

	length := defLen
	if len(inputLen) > 0 {
		length = inputLen[0]
	}

	input := RandomSlice(length)

	for sfname, sf := range sfm {
		t0 := time.Now()
		sf(input)
		tf := time.Since(t0)

		fmt.Printf("%s: %v\n", sfname, tf)
	}
}

func check(t *testing.T, tc testcase, got, clone []int) {
	if !reflect.DeepEqual(tc.expected, got) {
		t.Errorf("incorrect sort (%s): expected %v, got %v\n", tc.description, tc.expected, got)
	}

	if !reflect.DeepEqual(tc.input, clone) {
		t.Error("input slice should not be altered")
	}
}

func TestAll(t *testing.T) {
	for _, a := range sortingAlgorithms {
		t.Run(a.name, a.testFunc)
	}
}

func BenchmarkAll(b *testing.B) {
	for _, a := range sortingAlgorithms {
		b.Run(a.name, a.benchFunc)
	}
}
