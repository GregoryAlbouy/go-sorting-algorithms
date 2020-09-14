package algorithms

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	runTest(t, testSortingAlgorithm{"BubbleSort", BubbleSort, nil, nil})
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(RandomSlice(benchsize))
	}
}
