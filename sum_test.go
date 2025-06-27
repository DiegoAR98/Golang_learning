package main

import (
	"math/rand"
	"testing"
)

// helper to get the sequential sum
func sequentialSum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func TestConcurrentSumVariousSplits(t *testing.T) {
	cases := []struct {
		nums         []int
		goroutines   int
	}{
		{[]int{}, 1},
		{[]int{42}, 4},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{1, 2, 3, 4, 5}, 3},
		{rand.Perm(100), 8},   // uneven split
		{rand.Perm(101), 10},  // another uneven
	}

	for _, c := range cases {
		want := sequentialSum(c.nums)
		got := ConcurrentSum(c.nums, c.goroutines)
		if got != want {
			t.Errorf("ConcurrentSum(%d elems, %d gr) = %d; want %d",
				len(c.nums), c.goroutines, got, want)
		}
	}
}

func TestConcurrentSumRaceSafety(t *testing.T) {
	// run under `go test -race` to verify no data races
	// a moderately large random slice
	nums := make([]int, 1_000)
	for i := range nums {
		nums[i] = rand.Intn(100)
	}

	want := sequentialSum(nums)
	if got := ConcurrentSum(nums, 10); got != want {
		t.Fatalf("race-safety test failed: got %d, want %d", got, want)
	}
}
