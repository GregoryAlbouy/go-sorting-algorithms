package sorting

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	runTest(t, testAlgorithm{"BubbleSort", BubbleSort, nil, nil})
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(RandomSlice(benchsize))
	}
}
