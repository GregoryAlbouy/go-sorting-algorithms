package sorting

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	clog "github.com/gregoryalbouy/go-custom-log"
)

type testcase struct {
	description string
	input       interface{}
	expected    interface{}
}

type testableAlgorithm struct {
	name      string
	sortFunc  sortFunc
	testFunc  func(*testing.T)
	benchFunc func(*testing.B)
}

var (
	testedAlgorithms = []testableAlgorithm{
		{"BubbleSort", BubbleSort, TestBubbleSort, BenchmarkBubbleSort},
		{"SelectionSort", SelectionSort, TestSelectionSort, BenchmarkSelectionSort},
		{"InsertionSort", InsertionSort, TestInsertionSort, BenchmarkInsertionSort},
		{"MergeSort", MergeSort, TestMergeSort, BenchmarkMergeSort},
		{"MergeSortConc", MergeSortConc, TestMergeSortConc, BenchmarkMergeSortConc},
		{"QuickSort", QuickSort, TestQuickSort, BenchmarkQuickSort},
		{"QuickSortConc", QuickSortConc, TestQuickSortConc, BenchmarkQuickSortConc},
		{"RadixSort", RadixSort, TestRadixSort, BenchmarkRadixSort},
		{"GoSort", GoSort, TestGoSort, BenchmarkGoSort},
	}

	benchsize = 10000
)

func runTest(t *testing.T, a testableAlgorithm) {
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
		if a.name == "RadixSort" && tc.description == "relative ints" {
			continue
		}
		in := tc.input.([]int)
		clone := make([]int, len(in))
		copy(clone, in)

		got := a.sortFunc(in)
		check(t, tc, got)
		checkIntegrity(t, in, clone)
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

func check(t *testing.T, tc testcase, got interface{}) {
	if !reflect.DeepEqual(got, tc.expected) {
		t.Errorf("%s: expected %v, got %v", tc.description, tc.expected, got)
	}
}

// checkIntegrity verifies an input array is not altered by a sorting function.
func checkIntegrity(t *testing.T, input, clone []int) {
	if !reflect.DeepEqual(input, clone) {
		t.Error("input slice should not be altered")
	}
}

func TestAll(t *testing.T) {
	for _, a := range testedAlgorithms {
		t.Run(a.name, a.testFunc)
	}
}

func BenchmarkAll(b *testing.B) {
	for _, a := range testedAlgorithms {
		b.Run(a.name, a.benchFunc)
	}
}
