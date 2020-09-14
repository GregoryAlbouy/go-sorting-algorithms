package algorithms

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	runTest(t, testSortingAlgorithm{"QuickSort", QuickSort, nil, nil})
}

func TestQuickSortConc(t *testing.T) {
	runTest(t, testSortingAlgorithm{"QuickSortConc", QuickSortConc, nil, nil})
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(RandomSlice(benchsize))
	}
}

func BenchmarkQuickSortConc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSortConc(RandomSlice(benchsize))
	}
}

func BenchmarkCompareQuickSort(b *testing.B) {
	b.Run("QuickSort", BenchmarkQuickSort)
	b.Run("QuickSortConc", BenchmarkQuickSortConc)
}
