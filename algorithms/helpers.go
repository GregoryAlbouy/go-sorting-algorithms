package algorithms

import (
	"log"
	"math/rand"
	"time"
)

// Swap swaps values from a slice
func Swap(arr []int, i, j int) {
	if i < 0 || j < 0 || i >= len(arr) || j >= len(arr) {
		log.Fatal("index out of range")
	}
	arr[i], arr[j] = arr[j], arr[i]
}

// RandomSlice generates a slice of n random numbers
func RandomSlice(n int) (arr []int) {
	randMax := 1000000
	randSrc := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		arr = append(arr, rand.New(randSrc).Intn(randMax))
	}

	return
}
