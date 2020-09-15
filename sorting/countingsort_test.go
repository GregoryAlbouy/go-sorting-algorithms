package sorting

import "testing"

func TestCountingSort(t *testing.T) {
	runTest(t, testAlgorithm{"CountingSort", CountingSort, nil, nil})
}

func BenchmarkCountingSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountingSort(RandomSlice(benchsize))
	}
}
