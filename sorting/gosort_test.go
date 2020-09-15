package sorting

import "testing"

func TestGoSort(t *testing.T) {
	runTest(t, testableAlgorithm{"GoSort", GoSort, nil, nil})
}

func BenchmarkGoSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoSort(RandomSlice(benchsize))
	}
}
