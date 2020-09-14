package algorithms

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	runTest(t, testableAlgorithm{"MergeSort", MergeSort, nil, nil})
}

func TestMergeSortConc(t *testing.T) {
	runTest(t, testableAlgorithm{"MergeSortConc", MergeSortConc, nil, nil})
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(RandomSlice(benchsize))
	}
}

func BenchmarkMergeSortConc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSortConc(RandomSlice(benchsize))
	}
}

func BenchmarkCompareMergeSort(b *testing.B) {
	b.Run("MergeSort", BenchmarkMergeSort)
	b.Run("MergeSortConc", BenchmarkMergeSortConc)
}
