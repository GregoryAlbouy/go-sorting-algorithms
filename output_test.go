package main

import "testing"

func TestJSON(t *testing.T) {
	fr := dummyFinalResult()

	if err := NewOutput("test.json").Write(fr); err != nil {
		t.Error(err)
	}
}

func TestCSV(t *testing.T) {
	fr := dummyFinalResult()

	if err := NewOutput("test.csv").Write(fr); err != nil {
		t.Error(err)
	}
}

func dummyFinalResult() FinalResult {
	trs := []AlgoResult{
		{"TestSort", []SizeResult{{100, s}, {1000, 3 * s}, {10000, 5 * s}}},
		{"TestSort2", []SizeResult{{100, 2 * s}, {1000, 4 * s}, {10000, 6 * s}}},
	}

	fr := FinalResult{
		TestResults:   trs,
		Sizes:         []int{100, 1000, 10000},
		ExecutionTime: totalDuration(trs),
	}

	return fr
}
