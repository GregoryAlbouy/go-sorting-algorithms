package sorting

import (
	"log"
	"math"
	"math/rand"
	"time"
)

// RandomSlice generates a slice of n random numbers.
func RandomSlice(n int) (arr []int) {
	randMax := 1000000
	randSrc := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		arr = append(arr, rand.New(randSrc).Intn(randMax))
	}

	return
}

// swap swaps values from a slice at given indexes.
func swap(arr []int, i, j int) {
	if i < 0 || j < 0 || i >= len(arr) || j >= len(arr) {
		log.Fatal("index out of range")
	}
	arr[i], arr[j] = arr[j], arr[i]
}

// abs returns the absolute value of a given int. Avoids float64 conversion
// steps from math.Abs().
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// powInt returns a power n as ints. Avoids float64 conversion steps
// from math.Pow().
func powInt(a, n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= a
	}
	return res
}

// max returns the greatest number of two ints. Avoids float64 conversion
// steps from math.Max()
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// minmax returns the minimum value and the maximum value of given slice.
func minmax(arr []int) (int, int) {
	n := len(arr)
	min, max := arr[0], arr[0]

	for i := 1; i < n; i++ {
		v := arr[i]
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// digitAt returns the digit at position i (starting from the right at i=0).
func digitAt(n, i int) int {
	return (abs(n) / powInt(10, i)) % 10
}

// digitCount returns the number of digits that constitute number n.
func digitCount(n int) int {
	if n == 0 {
		return 1
	}
	return int(math.Floor(math.Log10(math.Abs(float64(n)))) + 1)
}

// mostDigits returns the number of digits of the integer having the greatest
// number of digits from the input slice.
func mostDigits(ns []int) int {
	maxDigits := 0
	for _, n := range ns {
		maxDigits = max(digitCount(n), maxDigits)
	}
	return maxDigits
}
