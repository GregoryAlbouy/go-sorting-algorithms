package algorithms

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestSwap(t *testing.T) {
	arr := []int{0, 42, 0, -42}
	expected := []int{-42, 42, 0, 0}

	Swap(arr, 1, 3)
	Swap(arr, 0, 3)
	Swap(arr, 0, 1)

	if !reflect.DeepEqual(arr, expected) {
		log.Panicf("expected %v, got %v", expected, arr)
	}
}

func TestRandom(t *testing.T) {
	fmt.Println(RandomSlice(10))
}
