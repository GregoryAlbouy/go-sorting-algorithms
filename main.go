package main

import (
	"fmt"
	"log"
	"time"

	clog "github.com/gregoryalbouy/go-custom-log"
	alg "github.com/gregoryalbouy/go-sorting-algorithms/algorithms"
)

// SortFunc type
type SortFunc func([]int) []int

// SortingAlgorithm struct
type SortingAlgorithm struct {
	name     string
	function SortFunc
}

// AlgoResult struct
type AlgoResult struct {
	Name    string       `json:"algorithm_name"`
	Results []SizeResult `json:"algorithm_results"`
}

// SizeResult struct
type SizeResult struct {
	Size     int           `json:"size"`
	Duration time.Duration `json:"duration"`
}

// FinalResult struct
type FinalResult struct {
	Algorithms    []string      `json:"algorithms"`
	Sizes         []int         `json:"sizes"`
	ExecutionTime time.Duration `json:"execution_time"`
	TestResults   []AlgoResult  `json:"results"`
}

var (
	defAlgorithms = []SortingAlgorithm{
		{"BubbleSort", alg.BubbleSort},
		{"SelectionSort", alg.SelectionSort},
		{"InsertionSort", alg.InsertionSort},
		{"MergeSort", alg.MergeSort},
		{"MergeSortConc", alg.MergeSortConc},
		{"QuickSort", alg.QuickSort},
		{"QuickSortConc", alg.QuickSortConc},
	}

	defInput = []int{100, 1000, 10000}
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	algos, sizes, outputs, err := parseInput()
	if err != nil {
		return err
	}

	testResults := genTests(algos, sizes)
	finalResult := FinalResult{
		Algorithms:    algoNames(algos),
		Sizes:         sizes,
		ExecutionTime: totalDuration(testResults),
		TestResults:   testResults,
	}

	for _, o := range outputs {
		if err := NewOutput(o).Write(finalResult); err != nil {
			return err
		}
	}

	return nil
}

func genTests(algos []SortingAlgorithm, sizes []int) []AlgoResult {
	var results []AlgoResult
	totalT0 := time.Now()

	for _, algo := range algos {
		var algoResults []SizeResult
		algoT0 := time.Now()
		fmt.Printf("%s ", algo.name)

		for _, size := range sizes {
			input := alg.RandomSlice(size)
			sizeT0 := time.Now()
			algo.function(input)
			sizeT1 := time.Since(sizeT0)

			algoResults = append(algoResults, SizeResult{size, sizeT1})
			fmt.Printf("| %d(%v) ", size, sizeT1)
		}

		algoT1 := time.Since(algoT0)
		results = append(results, AlgoResult{algo.name, algoResults})
		fmt.Printf("\n%s %s (%v)\n", clog.Green("OK"), algo.name, algoT1)
	}

	totalT1 := time.Since(totalT0)
	fmt.Printf("total %v\n", totalT1)

	return results
}

// totalDuration func
func totalDuration(ars []AlgoResult) time.Duration {
	var ns int64

	for _, ar := range ars {
		for _, r := range ar.Results {
			ns += r.Duration.Nanoseconds()
		}
	}

	return time.Duration(ns)
}
