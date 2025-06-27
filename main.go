package main

import (
	"fmt"
	"sync"
)

// ConcurrentSum splits nums into up to numGoroutines chunks,
// sums each chunk in its own goroutine, and safely accumulates
// into total using a mutex.
func ConcurrentSum(nums []int, numGoroutines int) int {
	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		total int
	)
	n := len(nums)
	if numGoroutines < 1 {
		numGoroutines = 1
	}

	// ceil(n / numGoroutines)
	chunkSize := (n + numGoroutines - 1) / numGoroutines

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if start >= n {
			// no work for this "goroutine"
			wg.Done()
			continue
		}
		if end > n {
			end = n
		}

		// capture the slice for this goroutine
		sub := nums[start:end]

		go func() {
			defer wg.Done()
			partial := 0
			for _, v := range sub {
				partial += v
			}
			mu.Lock()
			total += partial
			mu.Unlock()
		}()
	}

	wg.Wait()
	return total
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Sum with 3 goroutines:", ConcurrentSum(nums, 3))
}
