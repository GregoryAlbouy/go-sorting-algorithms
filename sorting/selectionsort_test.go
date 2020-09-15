package sorting

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	runTest(t, testableAlgorithm{"SelectionSort", SelectionSort, nil, nil})
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort(RandomSlice(benchsize))
	}
}
