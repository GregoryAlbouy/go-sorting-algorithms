package sorting

import (
	"testing"
)

func TestRadixSort(t *testing.T) {
	runTest(t, testAlgorithm{"RadixSort", RadixSort, nil, nil})
}

func BenchmarkRadixSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RadixSort(RandomSlice(benchsize))
	}
}
