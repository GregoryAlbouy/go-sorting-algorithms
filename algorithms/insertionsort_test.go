package algorithms

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	runTest(t, testSortingAlgorithm{"InsertionSort", InsertionSort, nil, nil})
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertionSort(RandomSlice(benchsize))
	}
}
