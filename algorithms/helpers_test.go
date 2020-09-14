package algorithms

import (
	"testing"
)

func TestSwap(t *testing.T) {
	tc := testcase{"swap test", []int{0, 42, 0, -42}, []int{-42, 42, 0, 0}}
	in := tc.input.([]int)

	swap(in, 1, 3)
	swap(in, 0, 3)
	swap(in, 0, 1)

	got := in
	check(t, tc, got)
}

func TestDigitAt(t *testing.T) {
	testcases := []testcase{
		{"position 0", []int{6293, 0}, 3},
		{"last position", []int{1000, 3}, 1},
		{"zero", []int{0, 0}, 0},
		{"negative number", []int{-837, 1}, 3},
	}

	for _, tc := range testcases {
		in := tc.input.([]int)
		got := digitAt(in[0], in[1])
		check(t, tc, got)
	}
}

func TestDigitCount(t *testing.T) {
	testcases := []testcase{
		{"positive number", 12345, 5},
		{"one digit", 8, 1},
		{"zero", 0, 1},
		{"negative number", -837, 3},
	}

	for _, tc := range testcases {
		in := tc.input.(int)
		got := digitCount(in)
		check(t, tc, got)
	}
}
