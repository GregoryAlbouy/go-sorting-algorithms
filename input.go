package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/gregoryalbouy/go-sorting-algorithms/sorting"
)

func parseInput() (algos []sorting.Algorithm, sizes []int, outputs []string, err error) {
	algoStr := flag.String("a", "", "algorithms to compare")
	sizeStr := flag.String("s", "", "sizes of input arrays")
	output := flag.String("o", "results.csv", "output path")
	flag.Parse()

	algos, err = algoSlice(*algoStr)
	if err != nil {
		return
	}
	sizes, err = intSlice(*sizeStr)
	if err != nil {
		return
	}
	outputs = strings.Fields(*output)

	return
}

func algoSlice(input string) ([]sorting.Algorithm, error) {
	if input == "" {
		return defAlgorithms, nil
	}

	inputSlice := strings.Fields(input)
	output := []sorting.Algorithm{}

	for _, s := range inputSlice {
		algo, err := findAlgoByName(s)
		if err != nil {
			return nil, err
		}

		output = append(output, algo)
	}

	return output, nil
}

func intSlice(input string) ([]int, error) {
	if len(input) == 0 {
		return defInput, nil
	}

	inputSlice := strings.Fields(input)
	output := []int{}

	for _, s := range inputSlice {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		output = append(output, n)
	}

	sort.Ints(output)
	return output, nil
}

func algoNames(algos []sorting.Algorithm) (names []string) {
	for _, a := range algos {
		names = append(names, a.Name)
	}
	return
}

func findAlgoByName(name string) (sorting.Algorithm, error) {
	for _, algo := range defAlgorithms {
		if strings.ToLower(name) == strings.ToLower(algo.Name) {
			return algo, nil
		}
	}

	return sorting.Algorithm{}, fmt.Errorf("algorithm %s does not exist", name)
}
